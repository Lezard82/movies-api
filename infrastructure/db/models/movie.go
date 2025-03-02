package models

import (
	"time"

	"github.com/Lezard82/movies-api/infrastructure/utils"
	"github.com/Lezard82/movies-api/internal/domain"
)

type MovieModel struct {
	ID          int64 `gorm:"primaryKey"`
	Title       string
	Director    string
	ReleaseDate time.Time
	Cast        string `gorm:"type:TEXT"` // JSON string to store
	Genre       string
	Synopsis    string
	UserID      int64     `gorm:"not null;index"` // Foreign key
	User        UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func FromDomainMovie(movie domain.Movie) *MovieModel {
	cast, _ := utils.MarshalCast(movie.Cast)

	return &MovieModel{
		ID:          movie.ID,
		Title:       movie.Title,
		Director:    movie.Director,
		ReleaseDate: movie.ReleaseDate,
		Cast:        cast,
		Genre:       movie.Genre,
		Synopsis:    movie.Synopsis,
		UserID:      movie.UserID,
	}
}

func (m *MovieModel) ToDomain() *domain.Movie {
	cast, _ := utils.UnmarshalCast(m.Cast)

	return &domain.Movie{
		ID:          m.ID,
		Title:       m.Title,
		Director:    m.Director,
		ReleaseDate: m.ReleaseDate,
		Cast:        cast,
		Genre:       m.Genre,
		Synopsis:    m.Synopsis,
		UserID:      m.UserID,
	}
}
