package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getEnv("port", "8080")
	fmt.Printf("Starting server at port %s", port)
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"msg\": \"hi\"}")
}

func getEnv(key, defaultVal string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		env = defaultVal
	}

	return env
}
