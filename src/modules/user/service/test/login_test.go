package login_test

import (
	"context"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
	services "task-management-be/src/modules/user/service"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	reqUser := domain.LoginRequest{
		Email:    "farisarma@gmail.com",
		Password: "xxxxx",
	}
	ctx := context.TODO()

	t.Run("Negative test case if user not found", func(t *testing.T) {
		mockRepository := new(MockRepository)
		// initialize user service using mock
		s := services.NewUserService(mockRepository)

		mockRepository.On("FindOne").Return(domain.User{}, helper.NewAppError(helper.ErrUserNotFound))
		login, err := s.Login(ctx, reqUser)

		expectedErr := helper.NewAppError(helper.ErrUserNotFound)

		assert.Error(t, err)
		assert.Equal(t, domain.LoginResponse{}, login)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("Positive test case", func(t *testing.T) {
		mockRepository := new(MockRepository)
		// initialize user service using mock
		s := services.NewUserService(mockRepository)
		user := domain.User{
			Id:         1,
			Name:       "Faris",
			Email:      "farisarma@gmail.com",
			Password:   "xxxxx",
			Created_At: time.Now(),
			Updated_At: time.Now(),
		}

		claims := helper.JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "LOGIN_USER",
				ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Minute)},
			},
			Email: user.Email,
			Name:  user.Name,
		}
		token, _ := helper.GenerateToken(claims)

		mockRepository.On("FindOne").Return(user, nil)
		expectedUser := domain.LoginResponse{
			Email:       "farisarma@gmail.com",
			AccessToken: token,
		}

		login, _ := s.Login(ctx, reqUser)
		assert.Equal(t, expectedUser, login)
	})

	t.Run("Negative test case if transaction db wont start", func(t *testing.T) {
		mockRepository := new(MockRepository)
		// initialize user service using mock
		s := services.NewUserService(mockRepository)

		expectedErr := helper.NewAppError(helper.ErrInternalServerError)
		mockRepository.On("FindOne").Return(domain.User{}, expectedErr)
		login, err := s.Login(ctx, reqUser)

		assert.Error(t, err)
		assert.Equal(t, domain.LoginResponse{}, login)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("Negative test case if failed find data transaction", func(t *testing.T) {
		// mockRepo := new(MockRepository)
		// s := services.NewUserService(mockRepo)

	})
}
