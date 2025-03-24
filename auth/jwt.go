package auth

import (
	"errors"
	"os"
	"time"

	"github.com/andrewbatallones/api/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

// Attempts to validate the token and return the user model.
func ValidateUserJWT(conn *pgxpool.Pool, tokenString string) (*models.User, error) {
	key, ok := os.LookupEnv("JWT_SALT")
	if !ok {
		return nil, errors.New("unable to find salt")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("can't pull claims")
	}
	
	// First need to parse it to float64
	// https://pkg.go.dev/encoding/json@go1.17.5#Number
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("can't pull user_id from claims")
	}

	return models.FindUser(conn, int(userId))
}
