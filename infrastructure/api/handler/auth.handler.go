package handler

import (
	"net/http"

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

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var userDTO struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user := domain.User{
		Username: userDTO.Username,
		Password: userDTO.Password,
	}

	if err := h.userUseCase.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) LoginUser(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := h.userUseCase.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
