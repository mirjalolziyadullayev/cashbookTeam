package main

import (
	"cashbookTeam/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server running...")

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/user", handlers.UserHandler)
	http.HandleFunc("/card", handlers.CardHandler)
	http.HandleFunc("/transaction", handlers.TransactionHandler)

	http.ListenAndServe(":8080", nil)
}