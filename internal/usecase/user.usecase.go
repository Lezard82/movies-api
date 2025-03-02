package usecase

import (
	"errors"

	"github.com/Lezard82/movies-api/infrastructure/security"
	"github.com/Lezard82/movies-api/internal/domain"
	"github.com/Lezard82/movies-api/internal/repository"
)

type UserUseCase struct {
	Repo       repository.UserRepository
	Hasher     security.PasswordHasher
	JWTService security.JWTService
}

func NewUserUseCase(repo repository.UserRepository, hasher security.PasswordHasher, jwt security.JWTService) *UserUseCase {
	return &UserUseCase{Repo: repo, Hasher: hasher, JWTService: jwt}
}

func (uc *UserUseCase) GetUserByID(id int64) (*domain.User, error) {
	return uc.Repo.GetByID(id)
}

func (uc *UserUseCase) GetUserByUsername(username string) (*domain.User, error) {
	return uc.Repo.GetByUsername(username)
}

func (uc *UserUseCase) GetAllUsers() ([]domain.User, error) {
	return uc.Repo.GetAll()
}

func (uc *UserUseCase) RegisterUser(user *domain.User) error {
	hashedPassword, err := uc.Hasher.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return uc.Repo.Create(user)
}

func (uc *UserUseCase) Authenticate(username string, password string) (string, error) {
	user, err := uc.Repo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password - username")
	}

	if !uc.Hasher.CheckPassword(user.Password, password) {
		return "", errors.New("invalid username or password - password")
	}

	token, err := uc.JWTService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *UserUseCase) UpdateUser(user *domain.User) error {
	return uc.Repo.Update(user)
}

func (uc *UserUseCase) DeleteUser(id int64) error {
	return uc.Repo.Delete(id)
}
