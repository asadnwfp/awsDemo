package main

import (
	"fmt"
	"net/http"

	"github.com/asadnwfp/mux/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", routes.GetHealth).Methods("GET")
	r.HandleFunc("/user", routes.SetUser).Methods("POST")
	r.HandleFunc("/users", routes.GetUsers).Methods("GET")

	// http.Handle("/", r)
	fmt.Println("Listen and Serve on Port :3000")
	http.ListenAndServe(":3000", r)

}
