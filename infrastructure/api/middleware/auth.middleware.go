package middleware

import (
	"net/http"
	"strings"

	"github.com/Lezard82/movies-api/infrastructure/api/dto"
	"github.com/Lezard82/movies-api/infrastructure/security"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtService security.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			dto.WriteErrorResponse(c, http.StatusUnauthorized, "Authorization header required")
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) < 2 {
			dto.WriteErrorResponse(c, http.StatusUnauthorized, "Invalid token format")
			c.Abort()
			return
		}

		token, err := jwtService.ValidateToken(tokenString[1])
		if err != nil || !token.Valid {
			dto.WriteErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			dto.WriteErrorResponse(c, http.StatusUnauthorized, "Invalid token claims")
			c.Abort()
			return
		}

		userID := int64(claims["user_id"].(float64))
		c.Set("user_id", userID)

		c.Next()
	}
}
