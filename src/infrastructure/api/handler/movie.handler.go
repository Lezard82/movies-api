package handler

import (
	"net/http"

	"github.com/Lezard82/movies-api/src/infrastructure/api/dto"
	"github.com/Lezard82/movies-api/src/infrastructure/api/helpers"
	"github.com/Lezard82/movies-api/src/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieUseCase *usecase.MovieUseCase
}

func NewMovieHandler(movieUseCase *usecase.MovieUseCase) *MovieHandler {
	return &MovieHandler{movieUseCase: movieUseCase}
}

// @Summary Get movies
// @Description Retrieves movies based on filters
// @Tags movies
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param title query string false "Title of the movie"
// @Param genre query string false "Genre of the movie"
// @Param release_date query string false "Release date of the movie"
// @Success 200 {array} domain.Movie "Filtered movies"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /movies [get]
func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	filters := make(map[string]interface{})
	if title := c.Query("title"); title != "" {
		filters["title"] = title
	}

	if releaseDate := c.Query("release_date"); releaseDate != "" {
		filters["release_date"] = releaseDate
	}

	if genre := c.Query("genre"); genre != "" {
		filters["genre"] = genre
	}

	movies, err := h.movieUseCase.GetAllMovies(filters)
	if err != nil {
		dto.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, movies)
}

// @Summary Get movie
// @Description Retrieves a movie based on ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "Movie ID"
// @Success 200 {object} domain.Movie "Certain movie"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /movies/{id} [get]
func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	id, err := helpers.GetMovieID(c)
	if err != nil {
		return
	}

	movie, err := h.movieUseCase.GetMovieByID(id)
	if err != nil {
		dto.WriteErrorResponse(c, http.StatusNotFound, "Movie not found")
		return
	}

	c.JSON(http.StatusOK, movie)
}

// @Summary Create movie
// @Description Create a movie object with parameters
// @Tags movies
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param request body dto.Movie true "Movie data"
// @Success 201 {object} domain.Movie "A complete movie"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /movies [post]
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
		dto.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// @Summary Update movie
// @Description Update a movie object with parameters
// @Tags movies
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param request body dto.Movie true "Movie data"
// @Param id path string true "Movie ID"
// @Success 200 {object} domain.Movie "A complete movie"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /movies/{id} [put]
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
		dto.WriteErrorResponse(c, http.StatusNotFound, "Movie not found")
		return
	}

	existingMovie.ID = id
	if helpers.UnauthorizedMovie(existingMovie, userID) {
		dto.WriteErrorResponse(c, http.StatusForbidden, "You are not allowed to edit this movie")
		return
	}

	movie, err := helpers.GetMovieFromContext(c)
	if err != nil {
		return
	}

	movie.ID = existingMovie.ID
	movie.UserID = existingMovie.UserID

	if err := h.movieUseCase.UpdateMovie(movie); err != nil {
		dto.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, movie)
}

// @Summary Delete movie
// @Description Deletes a movie based on ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "Movie ID"
// @Success 200 {object} dto.MessageResponse "Movie deleted"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Failure 403 {object} dto.ErrorResponse "Forbidden"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /movies/{id} [delete]
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
		dto.WriteErrorResponse(c, http.StatusNotFound, "Movie not found")
		return
	}

	if helpers.UnauthorizedMovie(movie, userID) {
		dto.WriteErrorResponse(c, http.StatusForbidden, "You are not allowed to delete this movie")
		return
	}

	if err := h.movieUseCase.DeleteMovie(movie.ID); err != nil {
		dto.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Movie deleted"})
}
