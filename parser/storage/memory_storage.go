package storage

import (
	"errors"
	"log"
	"parser/parser/model"
	"sync"
)

type MemoryTransactionDAO struct {
}

var (
	// {address -> list of transactions, newer transaction comes first}
	addressTransactionsMap map[string][]model.Transaction

	// key is the address which subscribed
	subscriptionMap map[string]bool

	// latestBlockNumber
	latestBlockNumber string
)

func (m MemoryTransactionDAO) Init() {
	log.Println("Memory TransactionDAO initialise start")

	var initLock sync.Mutex
	initLock.Lock()
	defer initLock.Unlock()

	if addressTransactionsMap == nil {
		addressTransactionsMap = make(map[string][]model.Transaction)
	}

	if subscriptionMap == nil {
		subscriptionMap = make(map[string]bool)
	}
	latestBlockNumber = ""
	log.Println("Memory TransactionDAO initialise finished")
}

func (m MemoryTransactionDAO) GetLatestBlockNumber() string {
	return latestBlockNumber
}

func (m MemoryTransactionDAO) UpdateLatestBlockNumber(number string) error {
	var initLock sync.Mutex
	initLock.Lock()
	defer initLock.Unlock()

	latestBlockNumber = number
	return nil
}

func (m MemoryTransactionDAO) SubscribeByAddress(address string) error {
	subscriptionMap[address] = true
	if _, ok := addressTransactionsMap[address]; !ok {
		addressTransactionsMap[address] = []model.Transaction{}
	}
	return nil
}

func (m MemoryTransactionDAO) FindByAddress(address string) (error, []model.Transaction) {
	if _, ok := subscriptionMap[address]; !ok {
		return errors.New("address not subscribed notification"), nil
	} else {
		return nil, addressTransactionsMap[address]
	}
}

func (m MemoryTransactionDAO) Save(transactions ...model.Transaction) error {
	for _, t := range transactions {
		parseAndSaveTransaction(t, t.From)
		parseAndSaveTransaction(t, t.To)
	}
	return nil
}

func parseAndSaveTransaction(transaction model.Transaction, address string) {
	if _, ok := addressTransactionsMap[address]; !ok {
		addressTransactionsMap[address] = []model.Transaction{}
	}
	existingTransactions := addressTransactionsMap[address]

	storedTxHash := make(map[string]bool, len(existingTransactions))
	for _, t := range existingTransactions {
		storedTxHash[t.Hash] = true
	}
	if _, ok := storedTxHash[transaction.Hash]; !ok {
		//log.Printf("parseAndSaveTransaction for address: %s", address)
		addressTransactionsMap[address] = append(addressTransactionsMap[address], transaction)
	}
}
