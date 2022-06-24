package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//区块结构体
type Block struct {
	Index     int64  //区块号
	TimeStamp int64  //时间戳
	PrevHash  string //父哈希
	Hash      string //当前哈希

	Data string //区块数据
}

//计算哈希值
func CalHash(block Block) string {
	blockData := string(block.Index) + string(block.TimeStamp) + block.PrevHash + block.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

//生成新区块
func GenNewBlock(pb Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = pb.Index + 1
	newBlock.PrevHash = pb.Hash
	newBlock.TimeStamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = CalHash(newBlock)
	return newBlock
}

//生成创世区块
func GenGenesisBlock() Block {
	pb := Block{}
	pb.Index = -1
	pb.Hash = ""
	return GenNewBlock(pb, "genesis block!!!")
}
