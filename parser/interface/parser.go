package _interface

import (
	"parser/parser/model"
)

type Parser interface {

	// GetCurrentBlock last parsed block
	GetCurrentBlock() int

	// Subscribe add address to observer
	Subscribe(address string) bool

	// GetTransactions list of inbound or outbound transactions for an address
	GetTransactions(address string) []model.Transaction

	// FetchLatestTransaction fetch latest transaction
	FetchLatestTransaction() error
}
