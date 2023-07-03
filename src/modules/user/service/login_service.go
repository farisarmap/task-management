package services

import (
	"context"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (u *UserService) Login(ctx context.Context, request domain.LoginRequest) (domain.LoginResponse, error) {
	user, err := u.Repository.FindOne(ctx, "email", request.Email)
	if err != nil {
		return domain.LoginResponse{}, helper.NewAppError(err.Error())
	}
	claims := helper.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "LOGIN_USER",
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Minute)},
		},
		Email: user.Email,
		Name:  user.Name,
	}

	token, err := helper.GenerateToken(claims)
	if err != nil {
		return domain.LoginResponse{}, helper.NewAppError(err.Error())
	}
	response := domain.LoginResponse{
		Email:       user.Email,
		AccessToken: token,
	}
	return response, nil
}
