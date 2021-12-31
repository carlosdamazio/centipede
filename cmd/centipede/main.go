package main

import (
    block "centipede/pkg/block"
    "fmt"
)

func main() {
    blk := block.GenesisBlock("This is the start of the blockchain")
    fmt.Println(blk.Data)
    
    blk = blk.NewBlock("Hey there! This is a new block!")
    fmt.Println(blk.Data)

    block.PrintBlockchain(blk)
}
