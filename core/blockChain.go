package core

import (
	"fmt"
	"log"
)

//区块链结构体
type BlockChain struct {
	Blocks []*Block
}

//创建区块链
func NewBlockchain() *BlockChain {
	genesisBlock := GenGenesisBlock()
	blockChain := BlockChain{}
	blockChain.AppendNewBlock(&genesisBlock)
	return &blockChain
}

//打包数据到区块
func (bc *BlockChain) SendDataToBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenNewBlock(*prevBlock, data)
	bc.AppendNewBlock(&newBlock)
}

//追加新区块
func (bc *BlockChain) AppendNewBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

//验证新区块是否合法
func isValid(newBlock Block, prevBlock Block) bool {
	if newBlock.Index != prevBlock.Index+1 {
		return false
	}
	if newBlock.PrevHash != prevBlock.Hash {
		return false
	}
	if newBlock.Hash != CalHash(newBlock) {
		return false
	}
	return true
}

//获取区块链的内容
func (bc *BlockChain) PrintBlockChain() {
	for _, v := range bc.Blocks {
		fmt.Printf("BlockIndex : %d\n", v.Index)
		fmt.Printf("Timestamp : %d\n", v.TimeStamp)
		fmt.Printf("PrevHash : %s\n", v.PrevHash)
		fmt.Printf("Hash : %s\n", v.Hash)
		fmt.Printf("Data : %s\n", v.Data)
		fmt.Println()
	}
}
