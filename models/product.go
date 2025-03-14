package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	Id          *int
	UserId      *int
	Title       string
	Description string
	Price       int
	IsAvailable bool
}

func (p *Product) Create(conn *pgxpool.Pool) error {
	query := "INSERT INTO products (user_id, title, description, price, is_available) VALUES ($1, $2, $3, $4, $5)"

	_, err := conn.Exec(context.Background(), query, p.UserId, p.Title, p.Description, p.Price, p.IsAvailable)
	return err
}
