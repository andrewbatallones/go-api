package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/andrewbatallones/api/cache"
	"github.com/andrewbatallones/api/models"
	"github.com/andrewbatallones/api/utils"
)

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	// Attempt to get cache then, send it if it's available
	foundCache := cache.GetCache(r.URL.Path)
	if foundCache != nil {
		w.Header().Set("Cache-Control", foundCache.CacheControl)
		w.Header().Set("Content-Type", foundCache.ContentType)
		w.Header().Set("Content-Length", foundCache.ContentLength)
		fmt.Fprint(w, foundCache.Body)

		return
	}

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

	body := fmt.Sprintf("{\"products\": [%s]}", json)

	// Set Cache
	newCache := cache.Cache{
		ContentType:   w.Header().Get("Content-Type"),
		ContentLength: w.Header().Get("Content-Length"),
		Body:          body,
	}

	err = newCache.SetCache(r.URL.Path)
	if err != nil {
		fmt.Printf("error setting cache: %s", err)
	}
	w.Header().Set("Cache-Control", newCache.CacheControl)

	fmt.Fprint(w, body)
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
