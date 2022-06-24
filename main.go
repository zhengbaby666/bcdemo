package main

import (
	"go_code/bcdemo/core"
	"go_code/bcdemo/rpc"
)

func main() {
	rpc.Blockchain = core.NewBlockchain()
	rpc.Run()
}
