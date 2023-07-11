package login_test

import (
	"context"
	"fmt"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
	services "task-management-be/src/modules/user/service"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	reqUser := domain.LoginRequest{}
	ctx := context.TODO()

	user := domain.User{
		Id:         1,
		Name:       "Faris",
		Email:      "farisarma@gmail.com",
		Password:   "xxxxx",
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}

	testCases := []struct {
		name         string
		mockRepoFunc func(repository *MockRepository)
		expectedUser domain.LoginResponse
		expectedErr  error
	}{
		{
			name: "Negative test case if user not found",
			mockRepoFunc: func(repository *MockRepository) {
				repository.On("FindOne").Return(domain.User{}, helper.NewAppError(helper.ErrUserNotFound))
			},
			expectedUser: domain.LoginResponse{},
			expectedErr:  helper.NewAppError(helper.ErrUserNotFound),
		},
		{
			name: "Positive test case",
			mockRepoFunc: func(repository *MockRepository) {
				repository.On("FindOne").Return(user, nil)
			},
			expectedUser: GenerateUserResponse(user),
			expectedErr:  nil,
		},
		{
			name: "Negative test case if transaction db wont start",
			mockRepoFunc: func(repository *MockRepository) {
				repository.On("FindOne").Return(domain.User{}, helper.NewAppError(helper.ErrInternalServerError))
			},
			expectedUser: domain.LoginResponse{},
			expectedErr:  helper.NewAppError(helper.ErrInternalServerError),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := new(MockRepository)
			s := services.NewUserService(mockRepository)

			tc.mockRepoFunc(mockRepository)

			login, err := s.Login(ctx, reqUser)
			fmt.Println(err, "<< error")

			assert.Equal(t, tc.expectedUser, login)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

}

func GenerateUserResponse(user domain.User) domain.LoginResponse {
	claims := helper.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "LOGIN_USER",
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Minute)},
		},
		Email: user.Email,
		Name:  user.Name,
	}

	token, _ := helper.GenerateToken(claims)
	expectedUser := domain.LoginResponse{
		Email:       "farisarma@gmail.com",
		AccessToken: token,
	}
	return expectedUser
}
