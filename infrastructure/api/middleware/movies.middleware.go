package middleware

import (
	"net/http"

	"github.com/Lezard82/movies-api/internal/domain"
	"github.com/gin-gonic/gin"
)

func ValidateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie domain.Movie

		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie data"})
			c.Abort()
			return
		}

		if movie.Title == "" || movie.Director == "" || movie.Genre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title, Director, and Genre are required"})
			c.Abort()
			return
		}

		c.Set("movie", movie)
		c.Next()
	}
}
