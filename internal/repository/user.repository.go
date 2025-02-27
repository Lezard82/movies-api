package repository

import (
	"github.com/Lezard82/movies-api/internal/domain"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetByID(id int64) (*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id int64) error
}
