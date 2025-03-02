package router

import (
	"github.com/Lezard82/movies-api/infrastructure/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine, authHandler *handler.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.RegisterUser)
		auth.POST("/login", authHandler.LoginUser)
	}
}
