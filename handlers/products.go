package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/andrewbatallones/api/models"
	"github.com/andrewbatallones/api/utils"
)

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	conn, ok := utils.Connection()
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to get products\"]}")
		return
	}
	defer conn.Close()

	products, err := models.AllProducts(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "issue getting products: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to get products\"]}")
		return
	}

	json, err := json.Marshal(products)
	if err != nil {
		fmt.Fprintf(os.Stderr, "issue marshalling products: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to get products\"]}")
		return
	}

	fmt.Fprintf(w, "{\"products\": [%s]}", json)
}
