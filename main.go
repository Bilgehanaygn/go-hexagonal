package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bilgehanaygn/urun/internal/category/application"
	router "github.com/bilgehanaygn/urun/internal/category/infra/http"
	"github.com/bilgehanaygn/urun/internal/category/infra/http/controller"
	internalpg "github.com/bilgehanaygn/urun/internal/category/infra/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type api struct {
	addr string
}

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

	defer m.Close()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	api := &api{addr: ":" + port}
	r := chi.NewRouter()
	server := &http.Server{Addr: api.addr, Handler: r}

	gormDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	initializeCategoryHandlerChainAndRegister(r, gormDB)

	log.Printf("Listening on %v", port)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func initializeCategoryHandlerChainAndRegister(r *chi.Mux, db *gorm.DB) {
	cRepo := internalpg.NewGormCategoryRepository(db)
	cSvc := application.CategoryService{CategoryRepository: cRepo}
	cCtrl := controller.CategoryController{CategoryService: cSvc}
	router.Register(r, &cCtrl)
}
