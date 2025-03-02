package handler

import (
	"net/http"

	"github.com/Lezard82/movies-api/infrastructure/api/dto"
	"github.com/Lezard82/movies-api/internal/domain"
	"github.com/Lezard82/movies-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewAuthHandler(userUseCase *usecase.UserUseCase) *AuthHandler {
	return &AuthHandler{userUseCase: userUseCase}
}

// @Summary Register a user
// @Description Registers a new user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body dto.Auth true "User registration data"
// @Success 201 {object} dto.MessageResponse "Login successful"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	auth := dto.Auth{}

	if err := c.ShouldBindJSON(&auth); err != nil {
		dto.WriteErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	user := domain.User{
		Username: auth.Username,
		Password: auth.Password,
	}

	if err := h.userUseCase.RegisterUser(&user); err != nil {
		dto.WriteErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, dto.MessageResponse{Message: "User registered successfully"})
}

// @Summary User login
// @Description Logs with an existing user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body dto.Auth true "User login data"
// @Success 200 {object} dto.AuthResponse "Json Web Token"
// @Failure 400 {object} dto.ErrorResponse "Bad request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) LoginUser(c *gin.Context) {
	auth := dto.Auth{}

	if err := c.ShouldBindJSON(&auth); err != nil {
		dto.WriteErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	token, err := h.userUseCase.Authenticate(auth.Username, auth.Password)
	if err != nil {
		dto.WriteErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{JsonWebToken: token})
}
