package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zacharygilliom/InventoryManager/internal/handlers"
	"github.com/zacharygilliom/InventoryManager/internal/models"
)

func main() {
	// Connect to our database and initialize our tables
	db := models.Connect()
	defer db.Close()

	models.CreateCustomerTable(db)
	models.CreateInventoryTable(db)
	models.CreateOrderTable(db)
	models.CreateOrderItemsTable(db)

	// Test data to insert
	insertData := make(map[string]string)
	tableName := "customer"
	insertData["name"] = "Smith&Smith"
	insertData["street_number"] = "112"
	insertData["street_name"] = "Main Street"
	insertData["city"] = "Youngstown"
	insertData["state"] = "Ohio"
	insertData["sales_region"] = "Central"
	models.InsertDataToTable(db, insertData, tableName)

	r := mux.NewRouter()
	r.HandleFunc("/customer/{id}", handlers.CustomerID(db)).Methods("GET")
	r.HandleFunc("/inventory/{id}", handlers.InventoryID(db)).Methods("GET")
	r.HandleFunc("/inventory/add/{item}/{price}/{quantity}", handlers.InventoryAdd(db)).Methods("POST", "GET", "PUT")
	http.ListenAndServe(":8080", r)

}
