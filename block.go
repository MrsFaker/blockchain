package main

import (
	"crypto/sha256"
	"time"
	"bytes"
	"encoding/binary"
)

const GenesisData = "hello 2018 summer"

//区块
type Block struct {
	Version    uint64 //版本号
	PrevHash   []byte //前区块哈希值
	MerkelRoot []byte //这是一个哈希值，后面v5用到
	TimeStamp  uint64 //时间戳，从1970.1.1到现在的秒数
	Difficulty uint64 //通过这个数字，算出一个哈希值：0x00010000000xxx
	Nonce      uint64 // 这是我们要找的随机数，挖矿就找证书
	Hash       []byte //当前区块哈希值, 正常的区块不存在，我们为了方便放进来
	Data       []byte //数据本身，区块体，先用字符串表示，v4版本的时候会引用真正的交易结构
}

//第一区块
func GenesisBlock(data string, prevHash []byte) *Block {

	return NewBlock(data, prevHash)
}
func NewBlock(data string, prevHash []byte) *Block {

	block := Block{
		Version:    00,
		PrevHash:   prevHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 100,
		Nonce:      100,
		Data:       []byte(data),
	}
	block.SetHash()
	return &block
}

//计算当前区块hash
func (block *Block) SetHash() {

	//Version    uint64 //版本号
	//PrevHash   []byte //前区块哈希值
	//MerkelRoot []byte //这是一个哈希值，后面v5用到
	//TimeStamp  uint64 //时间戳，从1970.1.1到现在的秒数
	//Difficulty uint64 //通过这个数字，算出一个哈希值：0x00010000000xxx
	//Nonce      uint64 // 这是我们要找的随机数，挖矿就找证书
	//Hash       []byte //当前区块哈希值, 正常的区块不存在，我们为了方便放进来
	//Data       []byte //数据本身，区块体，先用字符串表示，v4版本的时候会引用真正的交易结构
	//var blockInfo []byte
	//blockInfo = append(blockInfo, Uint2Byte(block.Version)...)
	//blockInfo = append(blockInfo, block.PrevHash...)
	//blockInfo = append(blockInfo, block.MerkelRoot...)
	//blockInfo = append(blockInfo, Uint2Byte(block.TimeStamp)...)
	//blockInfo = append(blockInfo, Uint2Byte(block.Difficulty)...)
	//blockInfo = append(blockInfo, Uint2Byte(block.Nonce)...)
	//blockInfo = append(blockInfo, block.Data...)

	byteArray := [][]byte{
		Uint2Byte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint2Byte(block.TimeStamp),
		Uint2Byte(block.Difficulty),
		Uint2Byte(block.Nonce),
		block.Data,
	}

	blockInfo := bytes.Join(byteArray, []byte{})

	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
func Uint2Byte(num uint64) []byte {
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, &num)
	return buffer.Bytes()
}
