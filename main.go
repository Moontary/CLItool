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
	CoreMenu(laptops)
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

// CoreMenu menu for looking through json
func CoreMenu(l []Laptop) {
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
			ReadProducts(l)
		case 2:
			FilterProducts()
		case 3:
			Exit()
		}
	}
}

// ReadProducts looks into json for products
func ReadProducts(l []Laptop) {
	i := 0
	for {
		var option int
		fmt.Printf("Product:\nID_%d\nName: %s\nDescription: %s\nPrice: %d\nSalesPrice: %d\nFeatures: %v\n", l[i].ID, l[i].Name, l[i].Description, l[i].Price, l[i].SalesPrice, l[i].Features)
		fmt.Println("1. Next")
		fmt.Println("2. Previous")
		fmt.Println("3. Edit")
		fmt.Println("4. Back")
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}
		switch option {
		case 1:
			changeIndex(&i, 1, len(l))
		case 2:
			changeIndex(&i, -1, len(l))
		case 3:
			EditProduct()
		case 4:
			return
		}
	}
}

// changeIndex changes
func changeIndex(i *int, n int, length int) {
	*i += n
	if *i == length {
		*i = 0
	} else if *i == -1 {
		*i = length - 1
	}
}

// EditProduct edits one product in json
func EditProduct() {
	fmt.Println("Edit product works")
}

// FilterProducts filters products in json
func FilterProducts() {
	fmt.Println("Filter products works")
}

// Exit exits program
func Exit() {
	os.Exit(3)
}
