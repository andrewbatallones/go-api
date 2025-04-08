package models

import (
	"context"
	"fmt"

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
	query := "INSERT INTO products (user_id, title, description, price, is_available) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	return conn.QueryRow(context.Background(), query, p.UserId, p.Title, p.Description, p.Price, p.IsAvailable).Scan(&p.Id)
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

func FindProduct(conn *pgxpool.Pool, product_id int) (*Product, error) {
	var p Product
	query := fmt.Sprintf("SELECT * FROM products WHERE id = %d LIMIT 1", product_id)

	err := conn.QueryRow(context.Background(), query).Scan(&p.Id, &p.UserId, &p.Title, &p.Description, &p.Price, &p.IsAvailable)

	return &p, err
}

func (p *Product) Update(conn *pgxpool.Pool) error {
	updater := fmt.Sprintf("SET title = '%s', description = '%s', price = %d, is_available = '%v'", p.Title, p.Description, p.Price, p.IsAvailable)
	_, err := conn.Exec(context.Background(), fmt.Sprintf("UPDATE products %s;", updater))

	return err
}
