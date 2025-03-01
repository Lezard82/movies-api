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
	Cast        string `gorm:"type:TEXT"` // Almacenamos como JSON string
	Genre       string
	Synopsis    string
	UserID      int64
}

func (m *MovieModel) FromDomain(movie domain.Movie) error {
	cast, err := utils.MarshalCast(movie.Cast)
	if err != nil {
		return err
	}

	m.ID = movie.ID
	m.Title = movie.Title
	m.Director = movie.Director
	m.ReleaseDate = movie.ReleaseDate
	m.Cast = cast
	m.Genre = movie.Genre
	m.Synopsis = movie.Synopsis
	m.UserID = movie.UserID

	return nil
}

func (m *MovieModel) ToDomain() (domain.Movie, error) {
	cast, err := utils.UnmarshalCast(m.Cast)
	if err != nil {
		return domain.Movie{}, err
	}

	return domain.Movie{
		ID:          m.ID,
		Title:       m.Title,
		Director:    m.Director,
		ReleaseDate: m.ReleaseDate,
		Cast:        cast,
		Genre:       m.Genre,
		Synopsis:    m.Synopsis,
		UserID:      m.UserID,
	}, nil
}
