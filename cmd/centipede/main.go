package main

import (
    "centipede/pkg/block"
)

func main() {
    blk := block.GenesisBlock("This is the start of the blockchain")
    blk = blk.NewBlock("Hey there! This is a new block!")
    data := blk.SerializeBlockchain()
    blk = block.DeserializeBlockchain(data)
}
