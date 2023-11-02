package main

import (
	"log"
	"net/http"
	"parser/parser"
	"parser/parser/scheduler"
	"parser/parser/storage"
)

var (
	currentBlockHandler   = parser.CurrentBlockHandler
	subscribeHandler      = parser.SubscribeHandler
	getTransactionHandler = parser.GetTransactionHandler

	// switch to other DAO implementation here
	dao storage.TransactionDAO = storage.MemoryTransactionDAO{}
)

func main() {
	dao.Init()

	err := scheduler.RefreshBlockAndTransaction()
	if err != nil {
		log.Fatal("failed to initialize server, error: ", err)
	}

	// Register the handler function with the default HTTP server
	http.HandleFunc("/get-current-block", currentBlockHandler)
	http.HandleFunc("/subscribe", subscribeHandler)
	http.HandleFunc("/get-transactions", getTransactionHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
