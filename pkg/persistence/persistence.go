package persistence

import (
    "fmt"
    "os"

    "centipede/pkg/block"
)

func PersistBlockchain(lastBlock *block.Block) error {
    data, err := lastBlock.SerializeBlockchain()
    if err != nil {
        return err
    }

    file, err := os.Create("centipede.dat")
    if err != nil {
        fmt.Fprintf(os.Stderr, err.Error()) 
        return err
    }

    file.Write(data)
    err = file.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, err.Error())
        return err
    }

    return nil
}

func ReadBlockchain(path string) (*block.Block, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    
    return block.DeserializeBlockchain(data)
}
