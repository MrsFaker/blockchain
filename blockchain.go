package main

type BlockChain struct {
	blocks []*Block
}

//区块链
func NewBlockChain() *BlockChain {
	block := GenesisBlock(GenesisData, []byte{})

	bc := BlockChain{[]*Block{block}}
	return &bc
}

//向链中添加block
func (bc *BlockChain) AddBlock(data string) {
	bcLen := len(bc.blocks)
	lastBlock := bc.blocks[bcLen-1]

	lastBlockHash := lastBlock.Hash

	block := NewBlock(data, lastBlockHash)
	bc.blocks = append(bc.blocks, block)

}
