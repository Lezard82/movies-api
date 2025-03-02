package handler

import (
	"net/http"

	"github.com/Lezard82/movies-api/infrastructure/api/helpers"
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
	id, err := helpers.GetMovieID(c)
	if err != nil {
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
	movie, err := helpers.GetMovieFromContext(c)
	if err != nil {
		return
	}

	userID, err := helpers.GetUserID(c)
	if err != nil {
		return
	}

	movie.UserID = userID

	if err := h.movieUseCase.CreateMovie(movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	userID, err := helpers.GetUserID(c)
	if err != nil {
		return
	}

	id, err := helpers.GetMovieID(c)
	if err != nil {
		return
	}

	existingMovie, err := h.movieUseCase.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	existingMovie.ID = id
	if helpers.UnauthorizedMovie(existingMovie, userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to edit this movie"})
		return
	}

	movie, err := helpers.GetMovieFromContext(c)
	if err != nil {
		return
	}

	movie.ID = existingMovie.ID
	movie.UserID = existingMovie.UserID

	if err := h.movieUseCase.UpdateMovie(movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	userID, err := helpers.GetUserID(c)
	if err != nil {
		return
	}

	id, err := helpers.GetMovieID(c)
	if err != nil {
		return
	}

	movie, err := h.movieUseCase.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if helpers.UnauthorizedMovie(movie, userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this movie"})
		return
	}

	if err := h.movieUseCase.DeleteMovie(movie.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
