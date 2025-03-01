package router

import (
	"github.com/Lezard82/movies-api/infrastructure/api/handler"
	"github.com/Lezard82/movies-api/infrastructure/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupMoviesRoutes(r *gin.Engine, movieHandler *handler.MovieHandler) {
	moviesGroup := r.Group("/movies")
	{
		moviesGroup.GET("/", movieHandler.GetAllMovies)
		moviesGroup.POST("/", middleware.ValidateMovie(), movieHandler.CreateMovie)
		moviesGroup.GET("/:id", movieHandler.GetMovieByID)
		moviesGroup.PUT("/:id", middleware.ValidateMovie(), movieHandler.UpdateMovie)
		moviesGroup.DELETE("/:id", movieHandler.DeleteMovie)
	}
}
