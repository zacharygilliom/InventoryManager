package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zacharygilliom/InventoryManager/internal/models"
)

// ID returns a database query with the requested customer ID
func ID(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		models.GetData(db, vars["id"])
	}
}
