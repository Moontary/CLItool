package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
)

var db *sql.DB

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
	DbConnection()
	defer db.Close()
	//laptops := make([]Laptop, 10)
	//ReadLaptops("laptops.json", &laptops)
	//featuresMap := getFeatures(laptops)
	//insertFeaturesDB(featuresMap)
	//insertLaptopsDB(laptops)
	//insertLaptopFeaturesDB(laptops)
	CoreMenu()
}

// getFeatures assignes features from json as keys in map
func getFeatures(l []Laptop) map[string]int {
	var featuresMap = make(map[string]int)
	for _, lap := range l {
		for _, feature := range lap.Features {
			featuresMap[feature] = 0
		}
	}
	return featuresMap
}

// insertFeaturesDB moves features from json to postgres
func insertFeaturesDB(featuresMap map[string]int) {
	query := "INSERT INTO feature(value) VALUES ($1);"
	for key, _ := range featuresMap {
		_, err := db.Exec(query, key)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// insertLaptopsDB moves laptop data from json to postgres
func insertLaptopsDB(l []Laptop) {
	query := "INSERT INTO Laptop(id, name, description, price, sales_price) VALUES ($1, $2, $3, $4, $5);"
	for _, lapData := range l {
		_, err := db.Exec(query, lapData.ID, lapData.Name, lapData.Description, lapData.Price, lapData.SalesPrice)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// insertLaptopFeaturesDB connects data from two tables into middle one
func insertLaptopFeaturesDB(laptops []Laptop) {
	query := "SELECT id FROM feature WHERE value = $1;"
	query2 := "INSERT INTO laptop_feature(laptop_id, feature_id) VALUES ($1, $2);"

	for _, laptop := range laptops {
		for _, feature := range laptop.Features {
			var featureID int
			row := db.QueryRow(query, feature)
			err := row.Scan(&featureID)
			if err != nil {
				fmt.Println(err)
			}
			_, err = db.Exec(query2, laptop.ID, featureID)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// DbConnection Connects to Database
func DbConnection() {
	var err error
	dbconnStr := "user=postgres dbname=postgres password=password host=localhost port=25432 sslmode=disable"
	db, err = sql.Open("postgres", dbconnStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
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
func CoreMenu() {

	var point int
	for {
		fmt.Println("1. Read Products")
		fmt.Println("2. Filter Products")
		fmt.Println("3. Exit")
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
	cmdClear()
	var (
		i = 0
		l = make([]Laptop, 0)
	)
	sqlStatement := `SELECT id, name, description, price, sales_price FROM laptop;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var laptop Laptop
		err = rows.Scan(&laptop.ID, &laptop.Name, &laptop.Description, &laptop.Price, &laptop.SalesPrice)
		l = append(l, laptop)
	}
	for {
		var option int
		fmt.Printf("Product:\nID_%d\nName: %s\nDescription: %s\nPrice: %d\nSalesPrice: %d\nFeatures: %v\n", l[i].ID, l[i].Name, l[i].Description, l[i].Price, l[i].SalesPrice, l[i].Features)
		fmt.Println("1. Next")
		fmt.Println("2. Previous")
		fmt.Println("3. Edit")
		fmt.Println("4. Back")
		_, err = fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}

		switch option {
		case 1:
			cmdClear()
			changeIndex(&i, 1, len(l))
		case 2:
			cmdClear()
			changeIndex(&i, -1, len(l))
		case 3:
			cmdClear()
			EditProduct(&l[i])
		case 4:
			cmdClear()
			return
		}
	}
}

// changeIndex iterates index for items in ReadProducts
func changeIndex(i *int, n int, length int) {
	*i += n
	if *i == length {
		*i = 0
	} else if *i == -1 {
		*i = length - 1
	}
}

// EditProduct edits one product in json
func EditProduct(l *Laptop) {
	cmdClear()
	sqlStatement := `UPDATE laptop SET name = $2, description = $3, price = $4 WHERE id = $1;`
	for {
		var option int
		fmt.Printf("Product:\nID_%d\nName: %s\nDescription: %s\nPrice: %d\nSalesPrice: %d\nFeatures: %v\n", l.ID, l.Name, l.Description, l.Price, l.SalesPrice, l.Features)
		fmt.Println("1. Change name")
		fmt.Println("2. Change description")
		fmt.Println("3. Change price")
		fmt.Println("4. Back")
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err)
		}
		switch option {
		case 1:
			fmt.Scan(&l.Name)
		case 2:
			fmt.Scan(&l.Description)
		case 3:
			fmt.Scan(&l.Price)
		case 4:
			_, err = db.Exec(sqlStatement, l.ID, l.Name, l.Description, l.Price)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}
}

// FilterProducts filters products in json
func FilterProducts() {
	cmdClear()
	sqlStatement := "select * from laptop where price between $1 and $2"
	for {
		var minPrice int
		var maxPrice int
		fmt.Println("Min Price to sort")
		fmt.Scan(&minPrice)

		fmt.Println("Max Price to sort")
		fmt.Scan(&maxPrice)
		if maxPrice < minPrice {
			fmt.Println("Max price can't be higher then min price")
			continue
		}
		rows, err := db.Query(sqlStatement, minPrice, maxPrice)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			var laptop Laptop
			err = rows.Scan(&laptop.ID, &laptop.Name, &laptop.Description, &laptop.Price, &laptop.SalesPrice)
			fmt.Printf("ID_%d\nName: %s\nDescription: %s\nPrice: %d\nSalesPrice: %d\n", laptop.ID, laptop.Name, laptop.Description, laptop.Price, laptop.SalesPrice)
		}

		return
	}
}

func cmdClear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Exit exits program
func Exit() {
	os.Exit(3)
}
