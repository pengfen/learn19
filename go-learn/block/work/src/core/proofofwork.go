package core

import (
	"math"
	"math/big"
	"bytes"
	"fmt"
	"crypto/sha256"
)

var maxNonce = math.MaxInt64

const targetBits = 20

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
		)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte)  {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("%s", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		/*
		采用哈希算法
		不可逆 无法从一个哈希值恢复原始数据 哈希并不是加密
		唯一性 对于特定的数据 只能有一个哈希 并且这个哈希是唯一的
		防篡改 改变输入数据中的一个字节 导致输出一个完全不同的哈希
		*/
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash) // 打印哈希值 (前五个数为0)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}