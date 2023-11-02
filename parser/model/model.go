package model

import "fmt"

type Transaction struct {
	BlockHash            string `json:"blockHash"`
	BlockNumber          string `json:"blockNumber"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	Hash                 string `json:"hash"`
	Input                string `json:"input"`
	Nonce                string `json:"nonce"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	TransactionIndex     string `json:"transactionIndex"`
	Value                string `json:"value"`
	Type                 string `json:"type"`
	ChainId              string `json:"chainId"`
}

// String of DaxTransaction
func (transaction *Transaction) String() string {
	s := fmt.Sprintf("blockHash=%s", transaction.BlockHash) +
		fmt.Sprintf("blockNumber=%s", transaction.BlockNumber) +
		fmt.Sprintf("gas=%s", transaction.Gas) +
		fmt.Sprintf("gasPrice=%s", transaction.GasPrice) +
		fmt.Sprintf("maxFeePerGas=%s", transaction.MaxFeePerGas) +
		fmt.Sprintf("maxPriorityFeePerGas=%s", transaction.MaxPriorityFeePerGas) +
		fmt.Sprintf("hash=%s", transaction.Hash) +
		fmt.Sprintf("input=%s", transaction.Input) +
		fmt.Sprintf("nonce=%s", transaction.Nonce) +
		fmt.Sprintf("from=%s", transaction.From) +
		fmt.Sprintf("to=%s", transaction.To) +
		fmt.Sprintf("transactionIndex=%s", transaction.TransactionIndex) +
		fmt.Sprintf("value=%s", transaction.Value) +
		fmt.Sprintf("type=%s", transaction.Type) +
		fmt.Sprintf("chainId=%s", transaction.ChainId)
	return s
}
