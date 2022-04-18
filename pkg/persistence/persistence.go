package persistence

import (
	"fmt"
	"log"
	"os"

	"github.com/carlosdamazio/centipede/pkg/block"
)

func PersistBlockchain(lastBlock *block.Block) error {
    data, err := lastBlock.SerializeBlockchain()
    if err != nil {
        return err
    }

    file, err := os.Create("centipede.dat")
    if err != nil {
        return err
    }

    file.Write(data)
    err = file.Close()
    if err != nil {
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
