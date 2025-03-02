package router

import (
	"github.com/Lezard82/movies-api/infrastructure/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(movieHandler *handler.MovieHandler, authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	SetupMoviesRoutes(r, movieHandler)
	SetupAuthRouter(r, authHandler)

	return r
}
