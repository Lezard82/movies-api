package repository

import (
	"github.com/Lezard82/movies-api/infrastructure/db"
	"github.com/Lezard82/movies-api/internal/domain"
)

type MovieRepositoryImpl struct {
	database db.Database
}

func NewMovieRepository(database db.Database) *MovieRepositoryImpl {
	return &MovieRepositoryImpl{database: database}
}

func (r *MovieRepositoryImpl) GetAll() ([]domain.Movie, error) {
	var movies []domain.Movie
	err := r.database.Find(&movies)
	return movies, err
}

func (r *MovieRepositoryImpl) GetByID(id int64) (*domain.Movie, error) {
	var movie domain.Movie
	err := r.database.First(&movie, id)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepositoryImpl) Create(movie *domain.Movie) error {
	return r.database.Create(movie)
}

func (r *MovieRepositoryImpl) Update(movie *domain.Movie) error {
	return r.database.Save(movie)
}

func (r *MovieRepositoryImpl) Delete(id int64) error {
	return r.database.Delete(&domain.Movie{}, id)
}
