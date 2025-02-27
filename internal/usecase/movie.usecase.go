package usecase

import (
	"github.com/Lezard82/movies-api/internal/domain"
	"github.com/Lezard82/movies-api/internal/repository"
)

type MovieUseCase struct {
	Repo repository.MovieRepository
}

func NewMovieUseCase(repo repository.MovieRepository) *MovieUseCase {
	return &MovieUseCase{Repo: repo}
}

func (uc *MovieUseCase) GetAllMovies() ([]domain.Movie, error) {
	return uc.Repo.GetAll()
}

func (uc *MovieUseCase) CreateMovie(movie *domain.Movie) error {
	return uc.Repo.Create(movie)
}

func (uc *MovieUseCase) UpdateMovie(movie *domain.Movie) error {
	return uc.Repo.Update(movie)
}

func (uc *MovieUseCase) DeleteMovie(id int64) error {
	return uc.Repo.Delete(id)
}
