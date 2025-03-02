package repository

import (
	"errors"

	"github.com/Lezard82/movies-api/infrastructure/db"
	"github.com/Lezard82/movies-api/infrastructure/db/models"
	"github.com/Lezard82/movies-api/internal/domain"
)

type UserRepositoryImpl struct {
	database db.Database
}

func NewUserRepository(database db.Database) *UserRepositoryImpl {
	return &UserRepositoryImpl{database: database}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	userModel := models.FromDomainUser(*user)
	return r.database.Create(&userModel)
}

func (r *UserRepositoryImpl) GetByUsername(username string) (*domain.User, error) {
	var userModel models.UserModel
	if err := r.database.FirstByField(&userModel, "username", username); err != nil {
		return nil, errors.New("user not found")
	}

	return userModel.ToDomain(), nil
}

func (r *UserRepositoryImpl) Delete(id int64) error {
	return r.database.Delete(&models.UserModel{}, id)
}

func (r *UserRepositoryImpl) Update(user *domain.User) error {
	userModel := models.FromDomainUser(*user)

	if err := r.database.Save(userModel); err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) GetAll() ([]domain.User, error) {
	var userModels []models.UserModel
	if err := r.database.Find(&userModels); err != nil {
		return nil, err
	}

	var users []domain.User
	for _, userModel := range userModels {
		user := userModel.ToDomain()
		users = append(users, *user)
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetByID(id int64) (*domain.User, error) {
	var userModel models.UserModel
	if err := r.database.First(&userModel, id); err != nil {
		return nil, err
	}

	return userModel.ToDomain(), nil
}
