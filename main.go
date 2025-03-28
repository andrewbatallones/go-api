package main

import (
	"net/http"

	"github.com/andrewbatallones/api/handlers"
	"github.com/andrewbatallones/api/middleware"
	"github.com/andrewbatallones/api/server"
	"github.com/andrewbatallones/api/utils"
)

func main() {
	port := utils.GetEnv("port", "8080")

	mux := http.NewServeMux()
	server := server.NewServer(mux, port)

	// Middleware
	server.WithMiddlewareFunc(middleware.Log)

	// Main
	server.WithHandler("/", handlers.Index)
	server.WithHandler("/healthcheck", handlers.Healthcheck)

	// Products
	server.WithHandler("/api/products", handlers.ProductIndex)
	server.WithHandler("/api/products/{product_id}", handlers.ProductShow)

	// Users
	server.WithHandler("/api/users", handlers.UserCreate)
	server.WithHandler("/api/users/{user_id}", handlers.UserShow)

	// Sessions
	server.WithHandler("/api/sessions", handlers.Sessions)

	server.Serve()
}
