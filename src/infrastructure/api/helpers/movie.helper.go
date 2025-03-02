package helpers

import (
	"net/http"

	"github.com/Lezard82/movies-api/src/infrastructure/api/dto"
	"github.com/Lezard82/movies-api/src/infrastructure/utils"
	"github.com/Lezard82/movies-api/src/internal/domain"
	"github.com/gin-gonic/gin"
)

func UnauthorizedMovie(movie *domain.Movie, userID int64) bool {
	return movie.UserID != userID
}

func GetUserID(c *gin.Context) (int64, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		dto.WriteErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return 0, http.ErrNoCookie
	}
	return userID.(int64), nil
}

func GetMovieID(c *gin.Context) (int64, error) {
	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		dto.WriteErrorResponse(c, http.StatusUnauthorized, "Invalid ID")
		return 0, err
	}
	return id, nil
}

func GetMovieFromContext(c *gin.Context) (*domain.Movie, error) {
	moviePtr, exists := c.Get("movie")
	if !exists {
		dto.WriteErrorResponse(c, http.StatusInternalServerError, "Movie data not found")
		return nil, http.ErrBodyNotAllowed
	}

	movie, ok := moviePtr.(*domain.Movie)
	if !ok {
		dto.WriteErrorResponse(c, http.StatusInternalServerError, "Invalid movie data")
		return nil, http.ErrBodyNotAllowed
	}

	return movie, nil
}
