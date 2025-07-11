package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bilgehanaygn/urun/internal/api"
	"github.com/bilgehanaygn/urun/internal/category/application"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/request"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/response"
	internalpg "github.com/bilgehanaygn/urun/internal/category/infra/postgres"
	"github.com/bilgehanaygn/urun/internal/pkg/app"
	"github.com/google/uuid"

	productapplication "github.com/bilgehanaygn/urun/internal/product/application"
	productreq "github.com/bilgehanaygn/urun/internal/product/infra/http/request"
	productres "github.com/bilgehanaygn/urun/internal/product/infra/http/response"
	productpg "github.com/bilgehanaygn/urun/internal/product/infra/postgres"

	catalogapplication "github.com/bilgehanaygn/urun/internal/catalog/application"
	catalogreq "github.com/bilgehanaygn/urun/internal/catalog/infra/http/request"
	catalogres "github.com/bilgehanaygn/urun/internal/catalog/infra/http/response"
	catalogpg "github.com/bilgehanaygn/urun/internal/catalog/infra/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//OTEL TRACING

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	gormotel "gorm.io/plugin/opentelemetry/tracing"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	m, err := migrate.New(
		"file://db/migrations",
		dbUrl,
	)

	if err != nil {
		log.Fatalf("failed to initialize migrations: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	config, err := app.NewConfig()
	if err != nil {
		log.Fatalf("Failed to initialize app config %v", err)
	}

	r := chi.NewRouter()
	server, err := initHttpServer(config, r)

	if err != nil {
		log.Fatalf("Failed to initialize http server: %v", err)
	}

	gormDB := initializeGorm(dbUrl)

	if _, err := startTracing(); err != nil {
		log.Fatalf("tracing init failed: %v", err)
	}

	categoryCPort := internalpg.NewCategoryCommandPort(gormDB)
	categoryQPort := internalpg.NewCategoryQueryPort(gormDB)
	categoryCreateHandler := &application.CategoryCreateHandler{CategoryCPort: categoryCPort}
	categoryUpdateHandler := &application.CategoryUpdateHandler{CategoryCPort: categoryCPort}
	categoryGetHandler := &application.CategoryQueryHandler{CategoryQPort: categoryQPort}

	productCPort := productpg.NewProductCommandPort(gormDB)
	productQPort := productpg.NewProductQueryPort(gormDB)
	productCreateHandler := &productapplication.ProductCreateHandler{ProductCPort: productCPort}
	productUpdateStatusHandler := &productapplication.ProductUpdateStatusHandler{ProductCPort: productCPort}
	productGetHandler := &productapplication.ProductGetHandler{ProductQPort: productQPort}

	catalogCPort := catalogpg.NewCatalogCommandPort(gormDB)
	catalogQPort := catalogpg.NewCatalogQueryPort(gormDB)
	catalogCreateHandler := &catalogapplication.CatalogCreateHandler{CatalogCPort: catalogCPort}
	catalogGetHandler := &catalogapplication.CatalogGetHandler{CatalogQPort: catalogQPort}

	r.Route("/category", func(r chi.Router) {
		r.Post("/", api.MakeHTTPHandler[request.CategoryCreateRequest, response.CategoryCreateResponse](categoryCreateHandler))
		r.Put("/", api.MakeHTTPHandler[request.CategoryUpdateRequest, response.CategoryUpdateResponse](categoryUpdateHandler))
		r.Get("/{id}", api.MakeHTTPHandler[uuid.UUID, response.CategoryDetailDto](categoryGetHandler))
	})
	r.Route("/product", func(r chi.Router) {
		r.Post("/", api.MakeHTTPHandler[productreq.ProductCreateRequest, productres.ProductCreateResponse](productCreateHandler))
		r.Get("/{id}", api.MakeHTTPHandler[productreq.ProductGetRequest, productres.ProductDetailDto](productGetHandler))
		r.Patch("/{productId}/status", api.MakeHTTPHandler[productreq.ProductUpdateStatusRequest, productres.ProductUpdateStatusResponse](productUpdateStatusHandler))
	})
	r.Route("/catalog", func(r chi.Router) {
		r.Post("/", api.MakeHTTPHandler[catalogreq.CatalogCreateRequest, catalogres.CatalogCreateResponse](catalogCreateHandler))
		r.Get("/{id}", api.MakeHTTPHandler[catalogreq.CatalogGetRequest, catalogres.CatalogDetailDto](catalogGetHandler))
	})

	go func() {
		log.Printf("Listening on %v", port)
		err = server.ListenAndServe()
	}()

	gracefulShutdown(server, gormDB)
}

func gracefulShutdown(srv *http.Server, gormDB *gorm.DB) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("interruption signal received, shutting down server…")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Printf("graceful shutdown did not complete in 5s: %v", err)
		if err2 := srv.Close(); err2 != nil {
			log.Printf("error forcing server close: %v", err2)
		}
	} else {
		log.Println("server shut down gracefully")
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Printf("could not retrieve raw DB from GORM: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("error closing DB pool: %v", err)
	} else {
		log.Println("database connection pool closed")
	}
}

func initializeGorm(dbUrl string) *gorm.DB {
	gormDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := gormDB.Use(gormotel.NewPlugin()); err != nil {
        log.Fatalf("failed to use GORM opentelemetry plugin: %v", err)
    }

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetConnMaxLifetime(time.Hour)

	return gormDB
}

func initHttpServer(config *app.Config, r *chi.Mux) (*http.Server, error) {

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   config.App.CORS.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	handler := otelhttp.NewHandler(r, "urun-service")

	server := &http.Server{Addr: ":" + config.App.Port, Handler: handler}

	return server, nil
}

func startTracing() (*trace.TracerProvider, error) {
	headers := map[string]string{
		"content-type": "application/json",
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("localhost:4318"),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating new exporter: %w", err)
	}

	tracerprovider := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("urun-service"),
			),
		),
		trace.WithSampler(trace.AlwaysSample()),
	)

	otel.SetTracerProvider(tracerprovider)

	return tracerprovider, nil
}
