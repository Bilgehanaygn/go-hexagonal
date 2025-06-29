package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bilgehanaygn/urun/internal/category/application"
	router "github.com/bilgehanaygn/urun/internal/category/infra/inp/http"
	"github.com/bilgehanaygn/urun/internal/category/infra/inp/http/controller"
	"github.com/bilgehanaygn/urun/internal/category/infra/out/db"
	"github.com/go-chi/chi/v5"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
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
	r := chi.NewRouter()

	server := &http.Server{Addr: api.addr, Handler: r}

	// gormDb := gorm.DB{}


	initializeCategoryHandlerChainAndRegister(r)

	log.Printf("Listening on %v", port)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


func initializeCategoryHandlerChainAndRegister(r *chi.Mux){
	cRepo := db.NewMockCategoryRepository(nil)
	cSvc := application.CategoryService{CategoryRepository: cRepo}
	cCtrl := controller.CategoryController{CategoryService: cSvc}
	router.Register(r, &cCtrl)
}