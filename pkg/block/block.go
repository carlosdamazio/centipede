package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
    Index     int
    Data      string
    Hash      string
    PrevBlock *Block 
}

// Creates new block from the given head
func (b *Block) NewBlock(data string) *Block {
    block := Block{b.Index + 1, data, "", b}
    block.Hash = calculateHash(b)
    return &block
}

// Generates the genesis block
func GenesisBlock(data string) *Block {
    block := Block{0, data, "", nil}
    block.Hash = calculateHash(&block)
    return &block
}

// Traverses and print hashes of the blocks in the blockchain
func PrintBlockchain(head *Block) {
    var curr *Block = head

    for ; curr != nil; {
        fmt.Printf("%s -> ", curr.Hash)
        curr = curr.PrevBlock 
    }
    fmt.Println("end")
}

// Calculates hash from given block being generated
func calculateHash(b *Block) string {
    var rec string = ""
    if b.PrevBlock == nil {
        rec = fmt.Sprint(b.Index) + b.Data 
    } else {
        rec = fmt.Sprint(b.Index) + b.Data + b.PrevBlock.Hash
    }

    h := sha256.New()
    h.Write([]byte(rec))
    hash := h.Sum(nil)
    return hex.EncodeToString(hash)
}
