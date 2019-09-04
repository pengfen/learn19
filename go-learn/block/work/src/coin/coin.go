package main

import (
	"core"
	"fmt"
	"strconv"
)

func main() {
	/*
		Prev hash:
        Data: Genesis Block
	    Hash: 7fd7bf892ecdb5cb7495ba8cf82d41c8958cf6d40ba0a776c40f46dfce651559
	*/
	bc := core.NewBlockchain() // 初始化区块链 创建第一个区块(创始区块)

	/*
		Prev hash: 7fd7bf892ecdb5cb7495ba8cf82d41c8958cf6d40ba0a776c40f46dfce651559
        Data: Send 1 BTC to ricky
        Hash: 4dfebf489b435cd6a1721b7eb4a1edc1cf7e16bd64181c18b1e80ed5560b1fc3
	*/
	bc.AddBlock("Send 1 BTC to ricky") // 加入一个区块 发送一个比特币给ricky
	/*
		Prev hash: 4dfebf489b435cd6a1721b7eb4a1edc1cf7e16bd64181c18b1e80ed5560b1fc3
        Data: Send 2 more BTC to ricky
        Hash: 610e78f49a44033a30f56c83f374e8735433497aecbaac742015c22d745c6773
	*/
	bc.AddBlock("Send 2 more BTC to ricky") // 加入一个区块 发送更多比特币给ricky

	for _,  block := range bc.Blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash) // 前一个区块的哈希值
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := core.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}