package repository

import (
	"github.com/Lezard82/movies-api/src/infrastructure/db"
	"github.com/Lezard82/movies-api/src/infrastructure/db/models"
	"github.com/Lezard82/movies-api/src/internal/domain"
)

type MovieRepositoryImpl struct {
	database db.Database
}

func NewMovieRepository(database db.Database) *MovieRepositoryImpl {
	return &MovieRepositoryImpl{database: database}
}

func (r *MovieRepositoryImpl) GetAll(filters map[string]interface{}) ([]domain.Movie, error) {
	var movieModels []models.MovieModel

	if err := r.database.Find(&movieModels, filters); err != nil {
		return nil, err
	}

	var movies []domain.Movie
	for _, movieModel := range movieModels {
		movie := movieModel.ToDomain()
		movies = append(movies, *movie)
	}

	return movies, nil
}

func (r *MovieRepositoryImpl) GetByID(id int64) (*domain.Movie, error) {
	var movieModel models.MovieModel
	if err := r.database.First(&movieModel, id); err != nil {
		return nil, err
	}

	return movieModel.ToDomain(), nil
}

func (r *MovieRepositoryImpl) Create(movie *domain.Movie) error {
	movieModel := models.FromDomainMovie(*movie)

	if err := r.database.Create(movieModel); err != nil {
		return err
	}

	movie.ID = movieModel.ID

	return nil
}

func (r *MovieRepositoryImpl) Update(movie *domain.Movie) error {
	movieModel := models.FromDomainMovie(*movie)

	if err := r.database.Save(movieModel); err != nil {
		return err
	}

	return nil
}

func (r *MovieRepositoryImpl) Delete(id int64) error {
	return r.database.Delete(&models.MovieModel{}, id)
}

func (r *MovieRepositoryImpl) Exists(movie *domain.Movie, excludeID int64) (bool, error) {
	conditions := map[string]interface{}{
		"title":        movie.Title,
		"director":     movie.Director,
		"release_date": movie.ReleaseDate,
	}

	count, err := r.database.CountByFields(&models.MovieModel{}, conditions, excludeID)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
