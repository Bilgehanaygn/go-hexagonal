package main

import (
	"net/http"
	"os"

	"github.com/bilgehanaygn/urun/internal/category/application"
	"github.com/bilgehanaygn/urun/internal/category/infra/inp/http/controller"
	"github.com/bilgehanaygn/urun/internal/category/infra/out/db"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from api"))
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	// dbUrl := os.Getenv("DB_URL")

	// m, err := migrate.New(
	// 	"file://db/migrations",
	// 	dbUrl,
	// )

	// if err != nil {
		// log.Fatalf("failed to initialize migrations: %v", err)
	// }

	// this is what actually *runs* the migrations
	// defer m.Close()
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatalf("migration failed: %v", err)
	// }

	api := &api{addr: ":" + port}
	mux := http.NewServeMux()

	server := &http.Server{Addr: api.addr, Handler: mux}

	gormDb := gorm.DB{}
	categoryRepository := db.NewGormCategoryRepository(&gormDb)
	categoryService := application.CategoryCommandService{CategoryRepository:categoryRepository}
	categoryCommandController := controller.CategoryCommandController{CategoryCommandService: categoryService}

	mux.HandleFunc("POST /kategori",  categoryCommandController.HandleCreate)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
