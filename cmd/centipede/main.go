package main

import (
    "fmt"
    "os"

    "centipede/pkg/persistence"
)

func main() {
    blk, err := persistence.ReadBlockchain("centipede.dat")
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err.Error())
        os.Exit(1)
    }

    fmt.Println(blk.Data)
    fmt.Println(blk.Hash)
    fmt.Println(blk.Index)
}
