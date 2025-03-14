package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	Id          *int   `json:"id"`
	UserId      *int   `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	IsAvailable bool   `json:"is_available"`
}

// Attempts to save the current product to the database.
func (p *Product) Create(conn *pgxpool.Pool) error {
	query := "INSERT INTO products (user_id, title, description, price, is_available) VALUES ($1, $2, $3, $4, $5)"

	_, err := conn.Exec(context.Background(), query, p.UserId, p.Title, p.Description, p.Price, p.IsAvailable)
	return err
}

func AllProducts(conn *pgxpool.Pool) ([]Product, error) {
	var products []Product
	query := "SELECT * FROM products"

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []Product{}, err
	}

	for rows.Next() {
		var p Product
		if err = rows.Scan(&p.Id, &p.UserId, &p.Title, &p.Description, &p.Price, &p.IsAvailable); err != nil {
			return []Product{}, err
		}

		products = append(products, p)
	}

	return products, nil
}
