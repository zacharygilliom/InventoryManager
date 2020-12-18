package main

import (
	database "github.com/zacharygilliom/InventoryManager/internal/models"
)

func main() {
	// Connect to our database and initialize our tables

	db := database.Connect()
	defer db.Close()
	database.CreateCustomerTable(db)
	database.CreateOrderTable(db)

	insertData := make(map[string]string)
	tableName := "customers"
	insertData["name"] = "Smith&Smith"
	insertData["street_number"] = "112"
	insertData["street_name"] = "Main Street"
	insertData["city"] = "Youngstown"
	insertData["state"] = "Ohio"
	insertData["sales_region"] = "Central"
	database.InsertDataToTable(db, insertData, tableName)

}
