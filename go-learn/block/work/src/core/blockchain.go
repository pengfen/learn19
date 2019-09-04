package core

// 区块链
type Blockchain struct {
	Blocks []*Block
}

// 添加一个区块至链上
func (bc *Blockchain) AddBlock(data string)  {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// 创建全新链
func NewBlockchain() *Blockchain  {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}