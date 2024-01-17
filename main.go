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

	/*
	Synchronous

		Main thread 
		(process 1 8s
		process 2 4s
		process 3 6s
		process 4 2s)

		Multithread
		1 thread process 1
		2 thread process 2
		3 thread process 3
		4 thread process 4

	Asynchronous
		Main thread {
			process 1 8s
			process 2 4s
			process 3 6s
			process 4 2s
		}
	*/
}