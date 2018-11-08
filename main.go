package main

import (
	"fmt"
)

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock("夜曲")
	blockChain.AddBlock("给我一首歌的时间")

	for i, block := range blockChain.blocks {
		fmt.Println("\n*******************block height:", i, "***********************")
		fmt.Printf("PrevBlock:%x\n", block.PrevHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
	}

}
