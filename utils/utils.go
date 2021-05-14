package utils

import (
	"bufio"
	"container/list"
	"fmt"
	"os"

	"github.com/onrcayci/gocalculator/grammar"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	HandleError(err)
	return input
}

func InfixToPostFix(tokens []string) []string {
	var postfix []string
	stack := list.New()
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]
		if !grammar.IsOperator(currentToken) {
			postfix = append(postfix, currentToken)
		} else if currentToken == "(" {
			stack.PushFront(currentToken)
		} else if currentToken == ")" {
			for stack.Len() != 0 && stack.Front().Value.(string) != "(" {
				top := stack.Front()
				stack.Remove(top)
				postfix = append(postfix, top.Value.(string))
			}
			lpar := stack.Front()
			stack.Remove(lpar)
		} else {
			for stack.Len() != 0 && grammar.GetOperatorPrecedence(currentToken) <= grammar.GetOperatorPrecedence(stack.Front().Value.(string)) {
				top := stack.Front()
				stack.Remove(top)
				postfix = append(postfix, top.Value.(string))
			}
			stack.PushFront(currentToken)
		}
	}

	// pop all remaining elements from the stack
	for stack.Len() != 0 {
		top := stack.Front()
		stack.Remove(top)
		postfix = append(postfix, top.Value.(string))
	}

	return postfix
}

func ExecuteInput(tokens []string) {
	switch tokens[0] {
	case "exit":
		os.Exit(0)
	default:
		fmt.Println(tokens)
		postfix := InfixToPostFix(tokens)
		fmt.Println(postfix)
	}
}
