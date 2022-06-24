package rpc

import (
	"encoding/json"
	"go_code/bcdemo/core"
	"io"
	"net/http"
)

//定义全局区块链
var Blockchain *core.BlockChain

//提供http服务
func Run() {
	http.HandleFunc("/blockchain/get", BlockChainGetHandler)
	http.HandleFunc("/blockchain/write", BlockChainWriteHandler)
	http.ListenAndServe("localhost:8888", nil)
}

//get
func BlockChainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(Blockchain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

//write
func BlockChainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	Blockchain.SendDataToBlock(blockData)
	BlockChainGetHandler(w, r)
}
