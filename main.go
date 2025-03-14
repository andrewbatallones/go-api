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

	// Main
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/healthcheck", handlers.Healthcheck)

	// Products
	http.HandleFunc("/api/products", handlers.ProductIndex)

	fmt.Printf("Starting server at port %s", port)
	fmt.Println()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
