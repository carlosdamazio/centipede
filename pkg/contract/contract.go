package contract

import (
    "fmt"
    "errors"
    "time"
)

// Function to generate code to be executed
func Execute() error {
    return errors.New("not implemented") 
}

// Interface to implement required methods.
type SmartContracter interface {
    Program(args []string) error
}

// The contract itself
type SmartContract struct {
    Name string
    Owner string
    Timestamp string
    Program func(args *[]string)
    Args *[]string
}

func (sc *SmartContract) Execute() {
    // Call function along with its args
    sc.Program(sc.Args)
}

// Generate a Smart Contract
func New(name string,
         owner string,
         program func(args *[]string),
         args *[]string) *SmartContract {
    return &SmartContract{
        Name: name,
        Owner: owner,
        Timestamp: fmt.Sprintf("%d", time.Now().UnixNano()),
        Program: program,
        Args: args,
    } 
}

