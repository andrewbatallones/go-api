package models

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       *int   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string `json:"-"`
}

func (u *User) Create(conn *pgxpool.Pool) error {
	query := "INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)"

	passwordHash, err := u.encryptPassword()
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), query, u.Name, u.Email, passwordHash)
	return err
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) encryptPassword() (string, error) {
	if len(u.password) == 0 {
		return "", errors.New("no password provided")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
