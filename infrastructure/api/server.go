package infrastructure

import (
	"github.com/Lezard82/movies-api/infrastructure/api/handler"
	"github.com/Lezard82/movies-api/infrastructure/api/router"
	"github.com/Lezard82/movies-api/infrastructure/db"
	"github.com/Lezard82/movies-api/internal/repository"
	"github.com/Lezard82/movies-api/internal/usecase"
)

func StartServer() {
	// Inicializar la base de datos
	database := db.InitDB()

	// Inicializar repositorios
	movieRepo := repository.NewMovieRepository(database)
	userRepo := repository.NewUserRepository(database)

	// Inicializar casos de uso
	movieUseCase := usecase.NewMovieUseCase(movieRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)

	movieHandler := handler.NewMovieHandler(movieUseCase)

	// Configurar el router
	r := router.SetupRouter(movieHandler)

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
