package handler

import (
	"net/http"

	"github.com/Lezard82/movies-api/infrastructure/utils"
	"github.com/Lezard82/movies-api/internal/domain"
	"github.com/Lezard82/movies-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieUseCase *usecase.MovieUseCase
}

func NewMovieHandler(movieUseCase *usecase.MovieUseCase) *MovieHandler {
	return &MovieHandler{movieUseCase: movieUseCase}
}

func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	movies, err := h.movieUseCase.GetAllMovies()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	id, err := utils.ParseID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	movie, err := h.movieUseCase.GetMovieByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	moviePtr, exists := c.Get("movie")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Movie data not found"})
		return
	}

	movie, ok := moviePtr.(*domain.Movie)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid movie data"})
		return
	}

	if err := h.movieUseCase.CreateMovie(movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id, err := utils.ParseID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	moviePtr, exists := c.Get("movie")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Movie data not found"})
		return
	}

	movie, ok := moviePtr.(*domain.Movie)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid movie data"})
		return
	}

	movie.ID = id

	if err := h.movieUseCase.UpdateMovie(movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id, err := utils.ParseID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.movieUseCase.DeleteMovie(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
