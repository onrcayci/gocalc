package main

import (
	"fmt"
	"os"

	"github.com/onrcayci/gocalculator/node"
	"github.com/onrcayci/gocalculator/parser"
	"github.com/onrcayci/gocalculator/utils"
)

func main() {

	// infinite app loop
	for {
		fmt.Printf(">> ")
		input := utils.GetInput()
		tokens := parser.Tokenize(input)
		switch tokens[0] {
		case "exit":
			os.Exit(0)
		default:
			postfix := utils.InfixToPostFix(tokens)
			root := node.BuildExpressionTree(postfix)
			result, err := root.SolveExpressionTree()
			utils.HandleError(err)
			fmt.Println(result)
		}
	}
}
