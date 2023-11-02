package _interface

import (
	"encoding/json"
	"log"
	"parser/parser/constant"
	"parser/parser/model"
	"parser/parser/storage"
	"parser/parser/vendors"
	"strconv"
	"strings"
)

const (
	BlockNumberSearchLimit = 10
)

var (
	// switch to other DAO implementation here
	dao storage.TransactionDAO = storage.MemoryTransactionDAO{}
)

type ETHParser struct {
}

func (p ETHParser) GetCurrentBlock() int {
	latestBlockNumber := dao.GetLatestBlockNumber()
	if latestBlockNumber == "" {
		log.Printf("GetCurrentBlock get response failed")
		return -1
	}

	// parse hex number to int
	intValue, err := strconv.ParseInt(strings.Replace(latestBlockNumber, constant.HexPrefix, "", -1), 16, 64)
	if err != nil {
		log.Printf("GetCurrentBlock parse block number failed: %v", err)
		return -1
	}
	return int(intValue)
}

func (p ETHParser) Subscribe(address string) bool {
	// basic parameter validation
	if !isValidRequest(address) {
		return false
	}

	err := dao.SubscribeByAddress(address)
	if err != nil {
		log.Printf("Subscribe failed for address: %s with error: %v", address, err)
	}
	return err == nil
}

func (p ETHParser) GetTransactions(address string) []model.Transaction {
	// basic parameter validation
	if !isValidRequest(address) {
		return nil
	}

	err, data := dao.FindByAddress(address)
	if err != nil {
		log.Printf("GetTransactions failed for address: %s with error: %v", address, err)
		return nil
	}
	return data
}

func (p ETHParser) FetchLatestTransaction() error {
	var param []interface{}
	var currentBlockResponse vendors.ETHCurrentBlockResponse
	err, responseBody := vendors.SendRequestToETHServer(vendors.GetCurrentBlockMethod, param)
	if err != nil {
		log.Printf("failed to fetchLatestTransaction with error: %v", err)
		return err
	}

	err = json.Unmarshal(responseBody, &currentBlockResponse)
	if err != nil {
		log.Printf("fetchLatestTransaction Unmarshal response failed with error: %+v", err)
		return err
	}

	if currentBlockResponse.Error.Message != "" {
		log.Printf("fetchLatestTransaction get response failed with error: %+v", currentBlockResponse.Error)
		return err
	}
	currentBlockNumber := currentBlockResponse.Result
	_ = dao.UpdateLatestBlockNumber(currentBlockNumber)
	parentBlockHash := ""

	i := 0
	for i < BlockNumberSearchLimit {
		if currentBlockNumber == "" && parentBlockHash == "" {
			break
		}
		err, transactionsOfThisBlock, nextParentBlockHash := fetchTransactionFromBlock(currentBlockNumber, parentBlockHash)

		// get block request failed, retry with same blockNumber and parentHash
		if err != nil {
			continue
		}
		err = dao.Save(transactionsOfThisBlock...)

		// store transaction failed, retry with same blockNumber and parentHash
		if err != nil {
			log.Printf("store transaction for block failed. blockNumber: %s, parentHash: %s", currentBlockNumber, parentBlockHash)
			continue
		}

		// proceed with parent block hash for next round
		parentBlockHash = nextParentBlockHash
		currentBlockNumber = ""
		i += 1
	}
	return nil
}

func fetchTransactionFromBlock(blockNumber string, parentBlockHash string) (error, []model.Transaction, string) {
	var err error
	var params []interface{}
	var transactionResponse vendors.ETHBlockTransactionResponse
	var responseBody []byte
	if blockNumber != "" {
		params = append(params, blockNumber)
		params = append(params, true)
		err, responseBody = vendors.SendRequestToETHServer(vendors.GetBlockByNumber, params)
	} else {
		params = append(params, parentBlockHash)
		params = append(params, true)
		err, responseBody = vendors.SendRequestToETHServer(vendors.GetBlockByHash, params)
	}
	if err != nil {
		log.Printf("get block failed blockNumber: %s, parentHash: %s", blockNumber, parentBlockHash)
		return err, nil, ""
	}
	err = json.Unmarshal(responseBody, &transactionResponse)
	if err != nil {
		log.Printf("fetchLatestTransaction Unmarshal response failed with error: %+v", err)
		return err, nil, ""
	}

	blockTransaction := transactionResponse.Result
	return nil, blockTransaction.Transactions, blockTransaction.ParentHash
}

func isValidRequest(address string) bool {
	return len(address) > 0 && strings.HasPrefix(address, constant.HexPrefix)
}
