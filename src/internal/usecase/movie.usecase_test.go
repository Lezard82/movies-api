package usecase

import (
	"testing"
	"time"

	"github.com/Lezard82/movies-api/src/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Repository Mock
type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) GetByID(id int64) (*domain.Movie, error) {
	args := m.Called(id)
	if movie, ok := args.Get(0).(*domain.Movie); ok {
		return movie, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMovieRepository) GetAll(filters map[string]interface{}) ([]domain.Movie, error) {
	args := m.Called(filters)
	return args.Get(0).([]domain.Movie), args.Error(1)
}

func (m *MockMovieRepository) Exists(movie *domain.Movie, id int64) (bool, error) {
	args := m.Called(movie, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockMovieRepository) Create(movie *domain.Movie) error {
	args := m.Called(movie)
	return args.Error(0)
}

func (m *MockMovieRepository) Update(movie *domain.Movie) error {
	args := m.Called(movie)
	return args.Error(0)
}

func (m *MockMovieRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

// === TESTS ===

func TestGetMovieByID(t *testing.T) {
	mockRepo := new(MockMovieRepository)
	usecase := NewMovieUseCase(mockRepo)

	expectedMovie := &domain.Movie{
		ID:          1,
		Title:       "Inception",
		Director:    "Christopher Nolan",
		ReleaseDate: time.Now(),
		Genre:       "Sci-Fi",
	}

	mockRepo.On("GetByID", int64(1)).Return(expectedMovie, nil)

	movie, err := usecase.GetMovieByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedMovie, movie)
	mockRepo.AssertExpectations(t)
}

func TestCreateMovie_Success(t *testing.T) {
	mockRepo := new(MockMovieRepository)
	usecase := NewMovieUseCase(mockRepo)

	newMovie := &domain.Movie{
		ID:          2,
		Title:       "The Dark Knight",
		Director:    "Christopher Nolan",
		ReleaseDate: time.Now(),
		Genre:       "Action",
	}

	mockRepo.On("Exists", newMovie, newMovie.ID).Return(false, nil)
	mockRepo.On("Create", newMovie).Return(nil)

	err := usecase.CreateMovie(newMovie)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateMovie_AlreadyExists(t *testing.T) {
	mockRepo := new(MockMovieRepository)
	usecase := NewMovieUseCase(mockRepo)

	existingMovie := &domain.Movie{
		ID:          3,
		Title:       "Interstellar",
		Director:    "Christopher Nolan",
		ReleaseDate: time.Now(),
		Genre:       "Sci-Fi",
	}

	mockRepo.On("Exists", existingMovie, existingMovie.ID).Return(true, nil)

	err := usecase.CreateMovie(existingMovie)

	assert.Error(t, err)
	assert.Equal(t, "a movie with the same title, director and release_date already exists", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestDeleteMovie(t *testing.T) {
	mockRepo := new(MockMovieRepository)
	usecase := NewMovieUseCase(mockRepo)

	mockRepo.On("Delete", int64(1)).Return(nil)

	err := usecase.DeleteMovie(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
