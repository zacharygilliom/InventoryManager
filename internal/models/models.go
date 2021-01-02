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
	Zip          string
	SalesRegion  string
}

// Customers ...
type Customers struct {
	List []Customer
}

// Order ...
type Order struct {
	ID         int
	CustomerID int
	Quantity   int
	Price      float64
}

// Item ...
type Item struct {
	ID   int
	name string
}

// Orders ...
type Orders struct {
	List []Order
}

// Inventory ...
type Inventory struct {
	ItemID   int
	Name     string
	Quantity int
	Price    string
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
	sqlStatement := `CREATE TABLE IF NOT EXISTS customer (
				customer_id 	serial			PRIMARY KEY,
				name 			varchar(40) 	NOT NULL,
				street_number	varchar(8) 		NOT NULL,
				street_name		varchar(15)		NOT NULL,
				city			varchar(15)		NOT NULL,
				state 			varchar(2)		NOT NULL,
				zip 			int				NOT NULL,
				sales_region	varchar(12)		NOT NULL
			)`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateOrderTable ...
func CreateOrderTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS orders (
				order_id 			serial			PRIMARY KEY,
				customer_id			int 			NOT NULL,
				quantity 			int 			NOT NULL,
				total_price			float 			NOT NULL,
				CONSTRAINT	fk_customer	FOREIGN KEY(customer_id)	REFERENCES customer(customer_id)
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
				item_name	varchar(15)	NOT NULL,
				quantity	int			NOT NULL,
				price 		float		NOT NULL
			)`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateOrderItemsTable ...
func CreateOrderItemsTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS orderitems (
				item_id int,
				order_id int,
				CONSTRAINT fk_item FOREIGN KEY(item_id) REFERENCES inventory(item_id),
				CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(order_id),
				PRIMARY KEY(item_id, order_id)
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
	sqlStatement := `SELECT * FROM customer WHERE customer_id =` + id
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
		err = rows.Scan(&i.ItemID, &i.Name, &i.Price, &i.Quantity)
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

// NewCustomer ...
func NewCustomer(db *sql.DB, data map[string]interface{}) {
	sqlStatement := `INSERT INTO customer (name, street_number, street_name, city, state, zip, sales_region)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(sqlStatement, data["name"], data["street_number"], data["street_name"], data["city"], data["state"], data["zip"], data["sales_region"])
	if err != nil {
		log.Fatal(err)
	}
}
