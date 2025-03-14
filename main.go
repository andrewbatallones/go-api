package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrewbatallones/api/utils"
)

func main() {
	port := utils.GetEnv("port", "8080")
	fmt.Printf("Starting server at port %s", port)
	http.HandleFunc("/", index)
	http.HandleFunc("/healthecheck", healthcheck)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"msg\": \"hi\"}")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"status\": \"ok\"}")
}
