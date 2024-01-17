package main

import (
	"cashbookTeam/handler"
	"fmt"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/", handler.HomeHandler).Methods("GET")
	r.HandleFunc("/user", handler.UserHandler).Methods("GET", "POST", "PUT", "DELETE")
	r.HandleFunc("/card", handler.CardHandler).Methods("GET", "POST", "PUT", "DELETE")
	r.HandleFunc("/transaction", handler.TransactionHandler).Methods("GET", "POST", "PUT", "DELETE")

	// Define the allowed origins, methods, and headers
	allowedOrigins := handlers.AllowedOrigins([]string{"http://127.0.0.1:5500", "http://localhost:5500", "https://mirjalolziyadullayev.github.io"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})

	// CORS
	corsHandler := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)
	fmt.Println("Server running...")

	http.Handle("/", corsHandler)
	http.ListenAndServe(":8080", nil)

}
