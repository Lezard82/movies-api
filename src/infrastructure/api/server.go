package api

import (
	"log"
	"net/http"
	"os"

	"github.com/Lezard82/movies-api/src/infrastructure/api/handler"
	"github.com/Lezard82/movies-api/src/infrastructure/api/router"
	"github.com/Lezard82/movies-api/src/infrastructure/db"
	"github.com/Lezard82/movies-api/src/infrastructure/repository"
	"github.com/Lezard82/movies-api/src/infrastructure/security"
	"github.com/Lezard82/movies-api/src/internal/usecase"
)

func StartServer() {
	database := db.InitDB()

	dbAdapter := db.NewGormDBAdapter(database)
	hasher := security.NewBcryptHasher()
	jwt := security.NewJWTService()

	movieRepo := repository.NewMovieRepository(dbAdapter)
	movieUseCase := usecase.NewMovieUseCase(movieRepo)
	movieHandler := handler.NewMovieHandler(movieUseCase)

	userRepo := repository.NewUserRepository(dbAdapter)
	userUseCase := usecase.NewUserUseCase(userRepo, hasher, jwt)
	authHandler := handler.NewAuthHandler(userUseCase)

	r := router.SetupRouter(movieHandler, authHandler, jwt)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
