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
