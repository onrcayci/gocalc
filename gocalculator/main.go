package main

import (
	"fmt"

	"github.com/onrcayci/gocalculator/parser"
	"github.com/onrcayci/gocalculator/utils"
)

func main() {

	// infinite app loop
	for {
		fmt.Printf(">> ")
		input := utils.GetInput()
		tokens := parser.Tokenize(input)
		utils.ExecuteInput(tokens)
	}
}
