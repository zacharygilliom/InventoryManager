package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	database "github.com/zacharygilliom/InventoryManager/internal/models"
)

// IDHandler ...
func IDHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cID := vars["id"]
		w.WriteHeader(http.StatusOK)
		database.GetData(db, cID)
	}
}

//HomeHandle ...
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World\n")
}

func main() {
	// Connect to our database and initialize our tables
	db := database.Connect()

	defer db.Close()
	//http.HandleFunc("/", HomeHandle)

	database.CreateCustomerTable(db)
	database.CreateInventoryTable(db)
	database.CreateOrderTable(db)

	// Test data to insert
	insertData := make(map[string]string)
	tableName := "customers"
	insertData["name"] = "Smith&Smith"
	insertData["street_number"] = "112"
	insertData["street_name"] = "Main Street"
	insertData["city"] = "Youngstown"
	insertData["state"] = "Ohio"
	insertData["sales_region"] = "Central"
	database.InsertDataToTable(db, insertData, tableName)

	r := mux.NewRouter()
	r.HandleFunc("/{id}", IDHandler(db)).Methods("GET")
	r.HandleFunc("/", HomeHandle).Methods("GET")
	http.ListenAndServe(":8080", r)

}
