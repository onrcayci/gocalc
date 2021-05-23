package main

import (
	"fmt"
	"os"

	"github.com/onrcayci/gocalc/node"
	"github.com/onrcayci/gocalc/parser"
	"github.com/onrcayci/gocalc/utils"
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
			root, err := node.BuildExpressionTree(postfix)
			utils.HandleError(err)
			result, err := root.SolveExpressionTree()
			utils.HandleError(err)
			fmt.Println(result)
		}
	}
}
