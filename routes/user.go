package routes

import (
	"encoding/json"
	"net/http"

	"github.com/asadnwfp/mux/data"
)

func SetUser(w http.ResponseWriter, r *http.Request) {
	var u data.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Add user to the in-memory store
	users := data.AddUser(u)

	// Return updated list as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.GetUsers())
}
