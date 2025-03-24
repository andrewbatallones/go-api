package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrewbatallones/api/handlers"
	"github.com/andrewbatallones/api/utils"
)

func main() {
	port := utils.GetEnv("port", "8080")

	mux := http.NewServeMux()

	// Main
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/healthcheck", handlers.Healthcheck)

	// Products
	mux.HandleFunc("/api/products", handlers.ProductIndex)
	mux.HandleFunc("/api/products/{product_id}", handlers.ProductShow)

	// Users
	mux.HandleFunc("/api/users", handlers.UserCreate)

	// Sessions
	mux.HandleFunc("/api/sessions", handlers.Sessions)

	fmt.Printf("Starting server at port %s", port)
	fmt.Println()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
