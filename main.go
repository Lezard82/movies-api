package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// Crear una instancia del router de Gin
	router := gin.Default()

	// Ruta de prueba (GET /ping)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Iniciar el servidor en el puerto 8080
	port := "8080"
	fmt.Println("Server running on http://localhost:" + port)
	router.Run(":" + port)
}