package usecase

import (
	"github.com/Lezard82/movies-api/internal/domain"
	"github.com/Lezard82/movies-api/internal/repository"
)

type UserUseCase struct {
	Repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{Repo: repo}
}

func (uc *UserUseCase) GetAllUsers() ([]domain.User, error) {
	return uc.Repo.GetAll()
}

func (uc *UserUseCase) CreateUser(user *domain.User) error {
	return uc.Repo.Create(user)
}

func (uc *UserUseCase) UpdateUser(user *domain.User) error {
	return uc.Repo.Update(user)
}

func (uc *UserUseCase) DeleteUser(id int64) error {
	return uc.Repo.Delete(id)
}
