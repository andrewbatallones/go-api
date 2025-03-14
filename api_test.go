package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrewbatallones/api/handlers"
)

func TestIndex(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	handlers.Index(resp, req)

	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get a successful response, got %d", resp.Result().StatusCode)
	}
}

func TestHealthcheck(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	resp := httptest.NewRecorder()

	handlers.Healthcheck(resp, req)

	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get a successful response, got %d", resp.Result().StatusCode)
	}
}

func TestProductIndex(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/api/products", nil)
	resp := httptest.NewRecorder()

	handlers.ProductIndex(resp, req)

	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get a successful response, got %d", resp.Result().StatusCode)
	}
}

func TestProductShow(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/api/products/1", nil)
	resp := httptest.NewRecorder()

	handlers.ProductIndex(resp, req)

	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("Did not get a successful response, got %d", resp.Result().StatusCode)
	}
}
