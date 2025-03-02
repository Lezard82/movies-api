package middleware

import (
	"net/http"

	"github.com/Lezard82/movies-api/infrastructure/api/dto"
	"github.com/Lezard82/movies-api/internal/domain"
	"github.com/gin-gonic/gin"
)

func ValidateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie domain.Movie

		if err := c.ShouldBindJSON(&movie); err != nil {
			dto.WriteErrorResponse(c, http.StatusBadRequest, "Invalid movie data")
			c.Abort()
			return
		}

		if movie.Title == "" || movie.Director == "" || movie.Genre == "" {
			dto.WriteErrorResponse(c, http.StatusBadRequest, "Title, Director, and Genre are required")
			c.Abort()
			return
		}

		if movie.ReleaseDate.IsZero() {
			dto.WriteErrorResponse(c, http.StatusBadRequest, "ReleaseDate is required and must be valid")
			c.Abort()
			return
		}

		c.Set("movie", &movie)
		c.Next()
	}
}
