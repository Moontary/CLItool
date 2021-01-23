package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Laptop stores info about laptops in json format
type Laptop struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	SalesPrice  int      `json:"sales_price"`
	Features    []string `json:"features"`
}

func main() {
	laptops := make([]Laptop, 10)
	ReadLaptops("laptops.json", &laptops)
	fmt.Println(laptops)
	Menu()
}

// ReadLaptops reads laptop from file and writes into slice
func ReadLaptops(namefile string, laptops *[]Laptop) {
	file, err := os.Open(namefile)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, laptops)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func Menu() {
	// menu for looking through json
	var point int
	for {
		fmt.Println("1. Read Products")
		fmt.Println("2. Filter Products")
		fmt.Println("3. Exit")
		// Takes users' input
		fmt.Scan(&point)
		fmt.Printf("Your choice %d\n", point)
		switch point {
		case 1:
			ReadProducts()
		case 2:
			FilterProducts()
		case 3:
			Exit()
		}
	}
}

// ReadProducts looks into json for products
func ReadProducts() {
	fmt.Println("Read products works")
}

// FilterProducts filters products from json
func FilterProducts() {
	fmt.Println("Filter products works")
}

// Exit exits program
func Exit() {
	os.Exit(3)
}
