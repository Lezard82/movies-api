package repository

import (
	"github.com/Lezard82/movies-api/infrastructure/db"
	"github.com/Lezard82/movies-api/infrastructure/db/models"
	"github.com/Lezard82/movies-api/internal/domain"
)

type MovieRepositoryImpl struct {
	database db.Database
}

func NewMovieRepository(database db.Database) *MovieRepositoryImpl {
	return &MovieRepositoryImpl{database: database}
}

func (r *MovieRepositoryImpl) GetAll() ([]domain.Movie, error) {
	var movieModels []models.MovieModel
	if err := r.database.Find(&movieModels); err != nil {
		return nil, err
	}

	var movies []domain.Movie
	for _, movieModel := range movieModels {
		movie, err := movieModel.ToDomain()
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *MovieRepositoryImpl) GetByID(id int64) (*domain.Movie, error) {
	var movieModel models.MovieModel
	if err := r.database.First(&movieModel, id); err != nil {
		return nil, err
	}

	movie, err := movieModel.ToDomain()
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *MovieRepositoryImpl) Create(movie *domain.Movie) error {
	movieModel := models.MovieModel{}

	if err := movieModel.FromDomain(*movie); err != nil {
		return err
	}

	if err := r.database.Create(movieModel); err != nil {
		return err
	}

	return nil
}

func (r *MovieRepositoryImpl) Update(movie *domain.Movie) error {
	movieModel := models.MovieModel{}

	if err := movieModel.FromDomain(*movie); err != nil {
		return err
	}

	if err := r.database.Save(movieModel); err != nil {
		return err
	}

	return nil
}

func (r *MovieRepositoryImpl) Delete(id int64) error {
	return r.database.Delete(&models.MovieModel{}, id)
}
