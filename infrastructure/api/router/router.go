package router

import (
	"github.com/Lezard82/movies-api/infrastructure/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(movieHandler *handler.MovieHandler) *gin.Engine {
	r := gin.Default()

	SetupMoviesRoutes(r, movieHandler)
	//users.SetpUserRoutes(r, userUseCase)

	return r
}
