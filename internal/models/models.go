package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "inventorymanager"
)

// Customer ...
type Customer struct {
	ID           int
	Name         string
	StreetNumber string
	StreetName   string
	City         string
	State        string
	SalesRegion  string
}

// Customers ...
type Customers struct {
	List []Customer
}

// Order ...
type Order struct {
	OrderID    int
	CustomerID int
	Quantity   int
	Product    string
	TotalPrice float64
}

// Orders ...
type Orders struct {
	List []Order
}

// Inventory ...
type Inventory struct {
	ID       int
	Price    string
	Name     string
	Quantity int
}

// Connect ...
func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host,
		port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected")
	return db
}

// CreateCustomerTable ...
func CreateCustomerTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS customers (
				customer_id 	serial			PRIMARY KEY,
				name 			varchar(40) 	NOT NULL,
				street_number	varchar(40) 	NOT NULL,
				street_name		varchar(40)		NOT NULL,
				city			varchar(40)		NOT NULL,
				state 			varchar(40)		NOT NULL,
				sales_region	varchar(40)		NOT NULL
			)`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateOrderTable ...
func CreateOrderTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS orders (
				order_id 				serial			PRIMARY KEY,
				customer_id				int 			NOT NULL,
				item_id					int				NOT NULL,
				item_price				float			NOT NULL,
				quantity 				int 			NOT NULL,
				item_name				varchar(15) 	NOT NULL,
				total_price				float 			NOT NULL,
				CONSTRAINT	fk_customer	FOREIGN KEY(customer_id)	REFERENCES customers(customer_id),
				CONSTRAINT 	fk_item FOREIGN KEY(item_id)			REFERENCES inventory(item_id)
			)`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateInventoryTable ...
func CreateInventoryTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS inventory (
				item_id		serial 		PRIMARY KEY,
				item_price	float		NOT NULL,
				item_name	varchar(15)	NOT NULL,
				quantity	int			NOT NULL 
	)`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertDataToTable ...
func InsertDataToTable(db *sql.DB, data map[string]string, tableName string) {
	sqlStatement := `INSERT INTO ` + tableName + `(name, street_number, street_name, city, state, sales_region)
			VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, data["name"], data["street_number"], data["street_name"], data["city"], data["state"], data["sales_region"])
	if err != nil {
		log.Fatal(err)
	}
}

// GetCustomerData ...
func GetCustomerData(db *sql.DB, id string) {
	sqlStatement := `SELECT * FROM customers WHERE customer_id =` + id
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var cs Customers

	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.ID, &c.Name, &c.StreetNumber, &c.StreetName, &c.City, &c.State, &c.SalesRegion)
		if err != nil {
			log.Fatal(err)
		}
		cs.List = append(cs.List, c)
		data, err := json.Marshal(c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", data)
		//fmt.Println(c)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

// GetInventoryData ...
func GetInventoryData(db *sql.DB, id string) {
	sqlStatement := `SELECT * FROM inventory WHERE item_id =` + id
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var i Inventory
		err = rows.Scan(&i.ID, &i.Name, &i.Price, &i.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		data, err := json.Marshal(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", data)
		//fmt.Println(c)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

//InsertDataToInventory ...
func InsertDataToInventory(db *sql.DB, data map[string]interface{}) {
	sqlStatement := `INSERT INTO inventory (item_name, item_price, quantity)
			VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlStatement, data["item_name"], data["item_price"], data["quantity"])
	if err != nil {
		log.Fatal(err)
	}
}
