package node

import (
	"container/list"
	"errors"
	"math"
	"strconv"

	"github.com/onrcayci/gocalc/grammar"
	"github.com/onrcayci/gocalc/utils"
)

type node struct {
	left  *node
	right *node
	value string
}

func NewNode(value string) *node {
	n := node{value: value}
	return &n
}

func BuildExpressionTree(tokens []string) *node {

	// instantiate an operand stack to store the operands of the operations
	operandStack := list.New()

	// loop through all of the input tokens
	for i := 0; i < len(tokens); i++ {
		currentToken := tokens[i]

		// if operand, push it into the operand stack
		if !grammar.IsOperator(currentToken) {
			operandNode := NewNode(currentToken)
			operandStack.PushFront(operandNode)
		} else {

			// initialize the operator node
			operatorNode := NewNode(currentToken)

			// pop the operand nodes from the stack
			operandNode1 := operandStack.Front()
			operandStack.Remove(operandNode1)
			operandNode2 := operandStack.Front()
			operandStack.Remove(operandNode2)

			// add operand nodes as children of the operator node
			operatorNode.right = operandNode1.Value.(*node)
			operatorNode.left = operandNode2.Value.(*node)

			// add the operator node into the stack in case it is the operand of another expression
			operandStack.PushFront(operatorNode)
		}
	}

	// get the root node from the stack
	expressionTree := operandStack.Front().Value.(*node)
	return expressionTree
}

func (n *node) SolveExpressionTree() (string, error) {
	if n != nil {
		if !grammar.IsOperator(n.value) {
			return n.value, nil
		}
		L, err := n.left.SolveExpressionTree()
		utils.HandleError(err)
		R, err := n.right.SolveExpressionTree()
		utils.HandleError(err)
		return Calculate(L, R, n.value), nil
	} else {
		return "", errors.New("cannot use nil expression to solve an equation")
	}
}

func Calculate(L string, R string, operator string) string {
	var result float64
	left, err := strconv.ParseFloat(L, 64)
	utils.HandleError(err)
	right, err := strconv.ParseFloat(R, 64)
	utils.HandleError(err)
	switch operator {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	case "^":
		result = math.Pow(left, right)
	}
	return strconv.FormatFloat(result, 'f', -1, 64)
}
