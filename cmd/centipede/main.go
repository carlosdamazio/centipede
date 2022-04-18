package main

import (
    "fmt"
    "os"

    "github.com/carlosdamazio/centipede/pkg/block"
    "github.com/carlosdamazio/centipede/pkg/contract"
    "github.com/carlosdamazio/centipede/pkg/persistence"
)

func main() {
    blk, err := persistence.ReadBlockchain("centipede.dat")
    if err != nil {
        if _, ok := err.(*os.PathError); !ok {
            fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err.Error())
            os.Exit(1)
        }
        blk = block.GenesisBlock("This is the genesis")
        con := contract.New("hello_world", "Carlos Damázio", `
            1 < 2
        `)
        blk = blk.NewBlock("this is a test", con)
        persistence.PersistBlockchain(blk)
        os.Exit(0)
    }

    blk.Contract.Execute()
}
