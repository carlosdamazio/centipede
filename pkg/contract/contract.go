package contract

import (
    "fmt"
    "errors"
    "go/token"
    "go/types"
    "os"
    "time"
)

// Function to generate code to be executed
func Execute() error {
    return errors.New("not implemented") 
}

// The contract itself
type SmartContract struct {
    Name string
    Owner string
    Timestamp string
    Program string
}

func (sc *SmartContract) Execute() {
    // Call function along with its args
    fs := token.NewFileSet()
    _, err := types.Eval(fs, nil, token.NoPos, sc.Program)
    if err != nil {
        fmt.Fprintf(os.Stderr, "couldn't evaluate program")
        os.Exit(1)
    }
}

// Generate a Smart Contract
func New(name, owner, program string) *SmartContract {
    return &SmartContract{
        Name: name,
        Owner: owner,
        Timestamp: fmt.Sprintf("%d", time.Now().UnixNano()),
        Program: program,
    } 
}

