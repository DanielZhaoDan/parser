package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"parser/parser/constant"
	_interface "parser/parser/interface"
	"parser/parser/model"
)

var (
	// switch to other parser implementation here
	parser = _interface.ETHParser{}
)

type CurrentBlockResponse struct {
	BlockNumber int `json:"blockNumber"`
}

type SubscribeResponse struct {
	Success bool `json:"success"`
}

type TransactionResponse struct {
	Transactions []model.Transaction `json:"transactions"`
}

func CurrentBlockHandler(w http.ResponseWriter, r *http.Request) {
	num := parser.GetCurrentBlock()
	assembleResponse(&CurrentBlockResponse{
		BlockNumber: num,
	}, w)
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	address := queryParams.Get("address")
	if address == "" {
		http.Error(w, errors.New("address parameter is required").Error(), http.StatusBadRequest)
	}
	res := parser.Subscribe(address)
	assembleResponse(&SubscribeResponse{
		Success: res,
	}, w)
}

func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	address := queryParams.Get("address")
	if address == "" {
		http.Error(w, errors.New("address parameter is required").Error(), http.StatusBadRequest)
	}
	res := parser.GetTransactions(address)
	assembleResponse(&TransactionResponse{
		Transactions: res,
	}, w)
}

func assembleResponse(data interface{}, w http.ResponseWriter) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set(constant.HeaderContentType, constant.JsonContentType)

	// Write the JSON response
	fmt.Fprintf(w, string(jsonData))
}
