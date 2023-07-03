package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Name  string `json:"name"`
}

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Now().Add(time.Hour * 24)
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func GenerateToken(claims JwtClaims) (string, error) {
	// claims := JwtClaims{
	// 	StandardClaims: jwt.RegisteredClaims{
	// 		Issuer:    APPLICATION_NAME,
	// 		ExpiresAt: &jwt.NumericDate{LOGIN_EXPIRATION_DURATION},
	// 	},
	// 	Email: email,
	// 	Name:  name,
	// }
	var pointingToken string
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	pointingToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		Logger.Warn().Msg("Failed to signed token")
		return pointingToken, fmt.Errorf(ErrJwtFailed)
	}
	return pointingToken, nil
}
