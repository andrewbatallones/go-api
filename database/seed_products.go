package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andrewbatallones/api/models"
	"github.com/andrewbatallones/api/utils"
	"github.com/go-faker/faker/v4"
)

func main() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to determine product size of %s", os.Args[1])
		os.Exit(1)
	}

	seedProducts(size)
}

func seedProducts(size int) {
	fmt.Println("Seeding products...")
	conn, ok := utils.Connection()
	if !ok {
		os.Exit(1)
	}
	defer conn.Close()

	for range size {
		var fp FakeProduct
		err := faker.FakeData(&fp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to fake product %v \n", err)
		}

		p := models.Product{
			Title:       fp.Title,
			Description: fp.Description,
			Price:       fp.Price,
			IsAvailable: true,
		}

		err = p.Create(conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot save product: %v\n", err)
		}
	}

	fmt.Println("Done!")
}

func printHelp() {
	fmt.Print(`This will seed in products into the app.
	
Beforehand, please be sure to create the database and pull in the structure.sql file.

when running the go command, pass in a number as an arguement:

Example:

	go run seed_products.go 1
	`)
}

type FakeProduct struct {
	Title       string `faker:"word"`
	Description string `faker:"sentence"`
	Price       int
}
