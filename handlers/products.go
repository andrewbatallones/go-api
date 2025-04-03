package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/andrewbatallones/api/models"
	"github.com/andrewbatallones/api/utils"
)

func ProductIndex(w http.ResponseWriter, r *http.Request) {
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

func ProductShow(w http.ResponseWriter, r *http.Request) {
	conn, ok := utils.Connection()
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to get product\"]}")
		return
	}
	defer conn.Close()

	product_id, err := strconv.Atoi(r.PathValue("product_id"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get product of product_id: %s", r.PathValue("product_id"))

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to get product\"]}")
		return
	}

	p, err := models.FindProduct(conn, product_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get product from database: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to get product\"]}")
		return
	}

	json, err := json.Marshal(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "issue marshalling product: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"errors\": [\"Unable to get product\"]}")
		return
	}

	fmt.Fprintf(w, "{\"product\": %s}", json)
}
