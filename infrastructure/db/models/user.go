package models

import "github.com/Lezard82/movies-api/internal/domain"

type UserModel struct {
	ID       int64 `gorm:"primaryKey"`
	Username string
	Password string
}

func (m *UserModel) FromDomain(user domain.User) {
	m.ID = user.ID
	m.Username = user.Username
	m.Password = user.Password
}

func (m *UserModel) ToDomain() domain.User {
	return domain.User{
		ID:       m.ID,
		Username: m.Username,
		Password: m.Password,
	}
}
