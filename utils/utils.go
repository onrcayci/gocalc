package utils

import (
	"bufio"
	"container/list"
	"os"

	"github.com/onrcayci/gocalc/grammar"
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
	for _, token := range tokens {
		if !grammar.IsOperator(token) {
			postfix = append(postfix, token)
		} else if token == "(" {
			stack.PushFront(token)
		} else if token == ")" {
			for stack.Len() != 0 && stack.Front().Value.(string) != "(" {
				top := stack.Front()
				stack.Remove(top)
				postfix = append(postfix, top.Value.(string))
			}
			lpar := stack.Front()
			stack.Remove(lpar)
		} else {
			for stack.Len() != 0 && grammar.GetOperatorPrecedence(token) <= grammar.GetOperatorPrecedence(stack.Front().Value.(string)) {
				top := stack.Front()
				stack.Remove(top)
				postfix = append(postfix, top.Value.(string))
			}
			stack.PushFront(token)
		}
	}
	for stack.Len() != 0 {
		top := stack.Front()
		stack.Remove(top)
		postfix = append(postfix, top.Value.(string))
	}
	return postfix
}
