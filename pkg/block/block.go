package block

import (
    "bytes"
	"crypto/sha256"
    "encoding/gob"
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

func DeserializeBlockchain(blockchainBytes []byte) (*Block, error) {
    buff := bytes.NewBuffer(blockchainBytes)
    dec := gob.NewDecoder(buff)

    var block Block
    err := dec.Decode(&block)
    if err != nil {
        return nil, err
    }

    return &block, nil
}

func (b *Block) SerializeBlockchain() ([]byte, error){
    var buffBytes []byte
    var res []byte
    buff := bytes.NewBuffer(buffBytes)
    enc := gob.NewEncoder(buff)
    curr := b

    for ; curr != nil ; {
        err := enc.Encode(*curr)
        if err != nil {
            return nil, err 
        }
        res = buff.Bytes()
        curr = curr.PrevBlock
    }

    return res, nil 
}
