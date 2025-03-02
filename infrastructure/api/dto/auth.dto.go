package dto

type Auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	JsonWebToken string `json:"jsonWebToken"`
}
