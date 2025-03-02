package dto

import "time"

type Movie struct {
	Title       string    `json:"title" binding:"required"`
	Director    string    `json:"director" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Cast        []string  `json:"cast"`
	Genre       string    `json:"genre"`
	Synopsis    string    `json:"synopsis"`
}
