package domain

import "time"

type Movie struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Director    string    `json:"director"`
	ReleaseDate time.Time `json:"release_date"`
	Cast        []string  `json:"cast"`
	Genre       string    `json:"genre"`
	Synopsis    string    `json:"synopsis"`
	UserID      int64     `json:"user_id"`
}
