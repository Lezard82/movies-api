package api

import (
	"log"
	"net/http"
	"os"

	"github.com/Lezard82/movies-api/infrastructure/api/handler"
	"github.com/Lezard82/movies-api/infrastructure/api/router"
	"github.com/Lezard82/movies-api/infrastructure/db"
	"github.com/Lezard82/movies-api/infrastructure/repository"
	"github.com/Lezard82/movies-api/internal/usecase"
)

func StartServer() {
	database := db.InitDB()

	dbAdapter := db.NewGormDBAdapter(database)
	movieRepo := repository.NewMovieRepository(dbAdapter)
	movieUseCase := usecase.NewMovieUseCase(movieRepo)
	movieHandler := handler.NewMovieHandler(movieUseCase)
	r := router.SetupRouter(movieHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
