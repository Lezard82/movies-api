package router

import (
	"github.com/Lezard82/movies-api/infrastructure/api/handler"
	"github.com/Lezard82/movies-api/infrastructure/security"
	"github.com/gin-gonic/gin"
)

func SetupRouter(movieHandler *handler.MovieHandler, authHandler *handler.AuthHandler, jwt security.JWTService) *gin.Engine {
	r := gin.Default()

	SetupMoviesRoutes(r, movieHandler, jwt)
	SetupAuthRouter(r, authHandler)
	SetupSwaggerRouter(r)

	return r
}
