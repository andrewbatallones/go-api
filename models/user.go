package models

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           *int   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	password     string `json:"-"`
	passwordHash string `json:"-"`
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

// Checks the password against the user.
func (u *User) CheckPassword(password string) bool {
	if len(password) == 0 {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(u.passwordHash), []byte(password)) == nil
}

func FindByUser(conn *pgxpool.Pool, params map[string]string) (*User, error) {
	whereClause := []string{}

	for field, value := range params {
		whereClause = append(whereClause, fmt.Sprintf("%s = '%s'", field, value))
	}

	var u User
	query := fmt.Sprintf("SELECT id, name, email, password_hash FROM users WHERE %s LIMIT 1", strings.Join(whereClause, " AND "))

	err := conn.QueryRow(context.Background(), query).Scan(&u.Id, &u.Name, &u.Email, &u.passwordHash)

	return &u, err
}
