package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zacharygilliom/InventoryManager/internal/models"
)

// InventoryID returns a database query with the requested customer ID
func InventoryID(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		models.GetInventoryData(db, vars["id"])
	}
}

// CustomerID returns a database query with the requested customer ID
func CustomerID(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		models.GetCustomerData(db, vars["id"])
	}
}

//InventoryAdd ...
func InventoryAdd(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		data := make(map[string]interface{})
		data["item_name"] = vars["item"]
		data["item_price"] = vars["price"]
		data["quantity"] = vars["quantity"]
		models.InsertDataToInventory(db, data)
	}
}

// AddCustomer ...
func AddCustomer(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		data := make(map[string]interface{})
		data["name"] = vars["name"]
		data["street_number"] = vars["street_number"]
		data["street_name"] = vars["street_name"]
		data["city"] = vars["city"]
		data["state"] = vars["state"]
		data["zip"] = vars["zip"]
		data["sales_region"] = vars["sales_region"]
		models.NewCustomer(db, data)
	}
}
