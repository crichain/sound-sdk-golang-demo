/**
* @Author:wjp
* @Date:2022/8/6 16:29
 */
package chainresp

//交易信息
type TransactionInfo struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RetCode     int `json:"retCode"`
		Transaction struct {
			Hash string `json:"hash"`
			Body struct {
				Nonce         int    `json:"nonce"`
				Address       string `json:"address"`
				Recipient     string `json:"recipient"`
				Amount        string `json:"amount"`
				InnerCodetype int    `json:"inner_codetype"`
				Timestamp     int64  `json:"timestamp"`
				Version       string `json:"version"`
				ChainId       int    `json:"chainId"`
				Deadline      int    `json:"deadline"`
			} `json:"body"`
			Signature string `json:"signature"`
			Node      struct {
				Nid     string `json:"nid"`
				Address string `json:"address"`
			} `json:"node"`
			Accepttimestamp int64 `json:"accepttimestamp"`
			Droptimestamp   int   `json:"droptimestamp"`
		} `json:"transaction"`
		Status struct {
			Status int `json:"status"`
			Height int `json:"height"`
		} `json:"status"`
		BlockHash string `json:"blockHash"`
	} `json:"data"`
}

//交易hash
type TxData struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RetCode int    `json:"retCode"`
		Hash    string `json:"hash"`
	} `json:"data"`
}

//用户信息
type Account struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RetCode int    `json:"retCode"`
		Address string `json:"address"`
		Nonce   int64  `json:"nonce"`
		Balance string `json:"balance"`
		Status  int    `json:"status"`
	} `json:"data"`
}
