package main

import (
	"fmt"
	"net/http"

	database "github.com/zacharygilliom/InventoryManager/internal/models"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}

func main() {
	// Connect to our database and initialize our tables

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)

	db := database.Connect()
	defer db.Close()
	database.CreateCustomerTable(db)
	database.CreateInventoryTable(db)
	database.CreateOrderTable(db)

	// Test data to insert
	/*
		insertData := make(map[string]string)
		tableName := "customers"
		insertData["name"] = "Smith&Smith"
		insertData["street_number"] = "112"
		insertData["street_name"] = "Main Street"
		insertData["city"] = "Youngstown"
		insertData["state"] = "Ohio"
		insertData["sales_region"] = "Central"
		database.InsertDataToTable(db, insertData, tableName)
	*/
}
