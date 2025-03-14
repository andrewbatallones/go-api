package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status := "ok"

	if !testConnection() {
		status = "database is unable to connect"
	}

	fmt.Fprintf(w, "{\"status\": \"%s\"}", status)
}

func testConnection() bool {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return false
	}
	defer conn.Close()

	var test string
	err = conn.QueryRow(context.Background(), "SELECT 'Testing'").Scan(&test)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false
	}

	return true
}
