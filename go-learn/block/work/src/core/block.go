package core

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Timestamp int64       // 区块创建时间
	Data []byte           // 区块包含数据
	PrevBlockHash []byte  // 前一个区块哈希值
	Hash []byte           // 区块自身的哈希值 用于校验区块数据有效

	Nonce int // 工作量证明
}

// 创建新区块并返回区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 10}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	block.SetHash()
	return block
}

// 设置hash值
func (b *Block) SetHash()  {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 创始区块
func NewGenesisBlock() *Block  {
	return NewBlock("Genesis Block", []byte{})
}