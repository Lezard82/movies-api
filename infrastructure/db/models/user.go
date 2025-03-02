package models

import (
	"github.com/Lezard82/movies-api/internal/domain"
)

type UserModel struct {
	ID       int64        `gorm:"primaryKey"`
	Username string       `gorm:"unique;not null"`
	Password string       `gorm:"not null"`
	Movies   []MovieModel `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

func FromDomainUser(user domain.User) *UserModel {
	return &UserModel{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}
}

func (m *UserModel) ToDomain() *domain.User {
	return &domain.User{
		ID:       m.ID,
		Username: m.Username,
		Password: m.Password,
	}
}
