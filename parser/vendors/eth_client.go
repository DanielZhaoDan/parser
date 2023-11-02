package vendors

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"parser/parser/constant"
	"parser/parser/model"
)

const (
	EthServerUrl          = "https://cloudflare-eth.com"
	RpcVersion            = "2.0"
	GetCurrentBlockMethod = "eth_blockNumber"
	GetBlockByNumber      = "eth_getBlockByNumber"
	GetBlockByHash        = "eth_getBlockByHash"
)

type ETHRequest struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int64         `json:"id"`
}

type ETHBasicResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Error   Error  `json:"error"`
	ID      int64  `json:"id"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ETHCurrentBlockResponse struct {
	ETHBasicResponse
	Result string `json:"result"`
}

type ETHBlockTransactionResponse struct {
	ETHBasicResponse
	Result ETHBlockTransaction `json:"result"`
}

type ETHBlockTransaction struct {
	Hash         string              `json:"hash"`
	ParentHash   string              `json:"parentHash"`
	Transactions []model.Transaction `json:"transactions"`
}

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

func SendRequestToETHServer(method string, param []interface{}) (error, []byte) {
	parameter := &ETHRequest{
		JsonRPC: RpcVersion,
		Method:  method,
		Params:  param,
		ID:      1,
	}
	body, _ := json.Marshal(parameter)

	// send post request and response check
	resp, err := http.Post(EthServerUrl, constant.JsonContentType, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("GetCurrentBlockFromETH failed with error: =%v", err)
		return err, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("GetCurrentBlockFromETH failed with httpCode: %d", resp.StatusCode)
		return err, nil
	}

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("GetCurrentBlockFromETH read response body failed: %v", err)
		return err, nil
	}
	return nil, responseBody
}
