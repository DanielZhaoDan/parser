package storage

import (
	"parser/parser/model"
)

type TransactionDAO interface {

	// Init initialize the DAO like create connections to database, initialize memory or cache
	Init()

	GetLatestBlockNumber() string

	UpdateLatestBlockNumber(number string) error

	// SubscribeByAddress subscribe transaction by address
	SubscribeByAddress(address string) error

	// FindByAddress find transactions by address
	FindByAddress(address string) (error, []model.Transaction)

	// Save store transactions
	Save(transactions ...model.Transaction) error
}
