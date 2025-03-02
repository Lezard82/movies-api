package repository

import (
	"github.com/Lezard82/movies-api/internal/domain"
)

type MovieRepository interface {
	GetAll(filters map[string]interface{}) ([]domain.Movie, error)
	GetByID(id int64) (*domain.Movie, error)
	Create(movie *domain.Movie) error
	Update(movie *domain.Movie) error
	Delete(id int64) error
	Exists(movie *domain.Movie, excludeID int64) (bool, error)
}
