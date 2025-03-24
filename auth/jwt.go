package auth

import (
	"errors"
	"os"
	"time"

	"github.com/andrewbatallones/api/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
}

type UserClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func JWTFromUser(u *models.User) (*JWT, error) {
	key, ok := os.LookupEnv("JWT_SALT")
	if !ok {
		return nil, errors.New("unable to find salt")
	}

	claims := UserClaims{
		*u.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "go-api",
			Subject:   u.Name,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	jwtStruct := JWT{
		Token:     token,
		TokenType: "Bearer",
		ExpiresIn: 2400,
	}

	return &jwtStruct, nil
}
