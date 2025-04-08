package models_test

import (
	"testing"

	"github.com/andrewbatallones/api/models"
	"github.com/andrewbatallones/api/utils"
)

func TestProductCreate(t *testing.T) {
	conn, ok := utils.Connection()
	if !ok {
		t.Error("unable to connect to database")
		return
	}
	defer conn.Close()

	newProduct := models.Product{
		Title:       "New Product",
		Description: "I'm a product?!",
		Price:       100,
		IsAvailable: true,
	}

	err := newProduct.Create(conn)
	if err != nil {
		t.Errorf("error when creating a product %s", err)
	}
}

func TestAllProducts(t *testing.T) {
	conn, ok := utils.Connection()
	if !ok {
		t.Error("unable to connect to database")
		return
	}
	defer conn.Close()

	_, err := models.AllProducts(conn)
	if err != nil {
		t.Errorf("error when creating a product %s", err)
	}
}

func TestProductUpdate(t *testing.T) {
	conn, ok := utils.Connection()
	if !ok {
		t.Error("unable to connect to database")
		return
	}
	defer conn.Close()

	product := models.Product{
		Title:       "New Product",
		Description: "I'm a product?!",
		Price:       100,
		IsAvailable: true,
	}

	err := product.Create(conn)
	if err != nil {
		t.Errorf("error when creating a product %s", err)
		return
	}

	product.Title = "Edited Product"
	product.Description = "huh..."
	product.Price = 10000

	err = product.Update(conn)
	if err != nil {
		t.Errorf("error when updating a product %s", err)
		return
	}

	// Get product and compare
	getProduct, err := models.FindProduct(conn, *product.Id)
	if err != nil {
		t.Errorf("error when getting a product %s", err)
		return
	}

	if compareProduct(*getProduct, product) {
		t.Errorf("the updated product does not match found product, expected %v, got %v", *getProduct, product)
	}
}

func compareProduct(p1, p2 models.Product) bool {
	return p1.Id == p2.Id &&
		p1.UserId == p2.UserId &&
		p1.Title == p2.Title &&
		p1.Description == p2.Description &&
		p1.Price == p2.Price &&
		p1.IsAvailable == p2.IsAvailable
}
