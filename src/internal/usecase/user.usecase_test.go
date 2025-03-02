package usecase_test

import (
	"errors"
	"testing"

	"github.com/Lezard82/movies-api/src/internal/domain"
	"github.com/Lezard82/movies-api/src/internal/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Repository Mock
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByID(id int64) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetAll(filters map[string]interface{}) ([]domain.User, error) {
	args := m.Called(filters)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockPasswordHasher struct {
	mock.Mock
}

func (m *MockPasswordHasher) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordHasher) CheckPassword(hashedPassword, password string) bool {
	args := m.Called(hashedPassword, password)
	return args.Bool(0)
}

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateToken(userID int64) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

func (m *MockJWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	args := m.Called(tokenString)
	return args.Get(0).(*jwt.Token), args.Error(1)
}

// === TESTS ===

func TestRegisterUser_InvalidPassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockJWT := new(MockJWTService)

	userUseCase := usecase.NewUserUseCase(mockRepo, mockHasher, mockJWT)

	user := &domain.User{
		Username: "testuser",
		Password: "123", // Contraseña demasiado corta
	}

	mockHasher.On("HashPassword", user.Password).Return("", nil)

	err := userUseCase.RegisterUser(user)
	assert.NotNil(t, err)
	assert.Equal(t, "insecure password", err.Error())
}

func TestRegisterUser_HashPasswordError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockJWT := new(MockJWTService)

	mockHasher.On("HashPassword", "Password123!").Return("", errors.New("hash error"))

	userUseCase := usecase.NewUserUseCase(mockRepo, mockHasher, mockJWT)

	user := &domain.User{
		Username: "testuser",
		Password: "Password123!",
	}

	err := userUseCase.RegisterUser(user)
	assert.NotNil(t, err)
	assert.Equal(t, "hash error", err.Error())
}

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockJWT := new(MockJWTService)

	mockHasher.On("HashPassword", "Password123!").Return("hashedpassword", nil)
	mockRepo.On("Create", mock.Anything).Return(nil)

	userUseCase := usecase.NewUserUseCase(mockRepo, mockHasher, mockJWT)

	user := &domain.User{
		Username: "testuser",
		Password: "Password123!",
	}

	err := userUseCase.RegisterUser(user)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
	mockHasher.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}

func TestAuthenticate_InvalidUsername(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockJWT := new(MockJWTService)

	mockRepo.On("GetByUsername", "testuser").Return((*domain.User)(nil), nil)

	userUseCase := usecase.NewUserUseCase(mockRepo, mockHasher, mockJWT)

	token, err := userUseCase.Authenticate("testuser", "Password123!")

	assert.Empty(t, token)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid username or password", err.Error())

	mockRepo.AssertExpectations(t)
	mockHasher.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}

func TestAuthenticate_InvalidPassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockJWT := new(MockJWTService)

	user := &domain.User{ID: 1, Username: "testuser", Password: "hashedpassword"}
	mockRepo.On("GetByUsername", "testuser").Return(user, nil)
	mockHasher.On("CheckPassword", "hashedpassword", "wrongpassword").Return(false)

	userUseCase := usecase.NewUserUseCase(mockRepo, mockHasher, mockJWT)

	token, err := userUseCase.Authenticate("testuser", "wrongpassword")
	assert.Empty(t, token)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid username or password - password", err.Error())
}

func TestAuthenticate_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockJWT := new(MockJWTService)

	user := &domain.User{ID: 1, Username: "testuser", Password: "hashedpassword"}
	mockRepo.On("GetByUsername", "testuser").Return(user, nil)

	mockHasher.On("CheckPassword", "hashedpassword", "Password123!").Return(true)

	mockJWT.On("GenerateToken", int64(1)).Return("valid_token", nil)

	userUseCase := usecase.NewUserUseCase(mockRepo, mockHasher, mockJWT)

	token, err := userUseCase.Authenticate("testuser", "Password123!")

	assert.Nil(t, err)
	assert.Equal(t, "valid_token", token)

	mockRepo.AssertExpectations(t)
	mockHasher.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}
