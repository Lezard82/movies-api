package router

import (
	"github.com/Lezard82/movies-api/infrastructure/api/handler"
	"github.com/Lezard82/movies-api/infrastructure/api/middleware"
	"github.com/Lezard82/movies-api/infrastructure/security"
	"github.com/gin-gonic/gin"
)

func SetupMoviesRoutes(r *gin.Engine, movieHandler *handler.MovieHandler, jwt security.JWTService) {
	moviesGroup := r.Group("/movies")

	moviesGroup.Use(middleware.AuthMiddleware(jwt))
	{
		moviesGroup.GET("/", movieHandler.GetAllMovies)
		moviesGroup.POST("/", middleware.ValidateMovie(), movieHandler.CreateMovie)
		moviesGroup.GET("/:id", movieHandler.GetMovieByID)
		moviesGroup.PUT("/:id", middleware.ValidateMovie(), movieHandler.UpdateMovie)
		moviesGroup.DELETE("/:id", movieHandler.DeleteMovie)
	}
}
