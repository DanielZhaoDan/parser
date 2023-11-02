package scheduler

import (
	"log"
	_interface "parser/parser/interface"
	"time"
)

const (
	// RefreshBlockIntervalInSecond average every 12 seconds there will be a new block in chain
	RefreshBlockIntervalInSecond = 12
)

var (
	requestParser _interface.Parser = _interface.ETHParser{}
)

func RefreshBlockAndTransaction() error {
	// Create a ticker that ticks every 10 seconds
	ticker := time.Tick(RefreshBlockIntervalInSecond * time.Second)

	err := doRefreshBlockAndTransaction()
	if err != nil {
		return err
	}
	// Run the code periodically
	go func() {
		for range ticker {
			doRefreshBlockAndTransaction()
		}
	}()
	return nil
}

func doRefreshBlockAndTransaction() error {
	log.Println("RefreshBlockAndTransaction start")
	err := requestParser.FetchLatestTransaction()
	if err != nil {
		log.Printf("RefrRefreshBlockAndTransaction with error: %v", err)
	}
	log.Println("RefreshBlockAndTransaction finish")
	return err
}
