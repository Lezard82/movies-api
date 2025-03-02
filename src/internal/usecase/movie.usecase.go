package usecase

import (
	"errors"

	"github.com/Lezard82/movies-api/src/internal/domain"
	"github.com/Lezard82/movies-api/src/internal/repository"
)

type MovieUseCase struct {
	Repo repository.MovieRepository
}

func NewMovieUseCase(repo repository.MovieRepository) *MovieUseCase {
	return &MovieUseCase{Repo: repo}
}

func (uc *MovieUseCase) GetMovieByID(id int64) (*domain.Movie, error) {
	return uc.Repo.GetByID(id)
}

func (uc *MovieUseCase) GetAllMovies(filters map[string]interface{}) ([]domain.Movie, error) {
	movies, err := uc.Repo.GetAll(filters)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (uc *MovieUseCase) CreateMovie(movie *domain.Movie) error {
	exists, err := uc.Repo.Exists(movie, movie.ID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("a movie with the same title, director and release_date already exists")
	}

	return uc.Repo.Create(movie)
}

func (uc *MovieUseCase) UpdateMovie(movie *domain.Movie) error {
	exists, err := uc.Repo.Exists(movie, movie.ID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("a movie with the same title, director and release_date already exists")
	}

	return uc.Repo.Update(movie)
}

func (uc *MovieUseCase) DeleteMovie(id int64) error {
	return uc.Repo.Delete(id)
}
