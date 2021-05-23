package node_test

import (
	"fmt"
	"testing"

	"github.com/onrcayci/gocalc/node"
)

var testTokens []string = []string{"1", "2", "+"}
var invalidTokens []string = []string{"1", "+"}
var nilTokens []string = make([]string, 1)

func TestBuildExpressionTree_ReturnsTree(t *testing.T) {
	tree, err := node.BuildExpressionTree(testTokens)
	testTree := tree.String()
	fmt.Println(testTree)
	if testTree != "1 + 2" && err != nil {
		t.Errorf("Fail to build expression tree for the postfix expression %v\n", testTokens)
	}
}

func TestBuildExpressionTree_ReturnsErr(t *testing.T) {
	tree, err := node.BuildExpressionTree(invalidTokens)
	if tree != nil && err == nil {
		t.Errorf("Fail to return an error for the invalid postfix expression %v\n", invalidTokens)
	}
}

func TestSolveExpressionTree_ReturnsTrue(t *testing.T) {
	tree, err := node.BuildExpressionTree(testTokens)
	if err != nil {
		t.Errorf("Fail to build expression tree for the postfix expression %v\n", testTokens)
	}
	result, err := tree.SolveExpressionTree()
	if result != "3" && err != nil {
		t.Errorf("Fail to solve expression tree for the postfix expression %v\n", testTokens)
	}
}

func TestSolveExpressionTree_ReturnsErr(t *testing.T) {
	tree, err := node.BuildExpressionTree(nilTokens)
	if tree != nil && err != nil {
		t.Errorf("Fail to build a nil expression tree\n")
	}
	result, err := tree.SolveExpressionTree()
	if result != "" && err == nil {
		t.Errorf("Fail to return an error for the nil expression tree\n")
	}
}

func TestString_ReturnsInfixExpression(t *testing.T) {
	tree, err := node.BuildExpressionTree(testTokens)
	if err != nil {
		t.Errorf("Fail to build expression tree for the postfix expression %v\n", testTokens)
	}
	treeString := tree.String()
	if treeString != "1 + 2" {
		t.Errorf("Fail to return the correct string representation of the expression tree\n")
	}
}

func TestString_ReturnsEmpty(t *testing.T) {
	tree, err := node.BuildExpressionTree(nilTokens)
	if err != nil {
		t.Errorf("Fail to build a nil expression tree\n")
	}
	treeString := tree.String()
	if treeString != "" {
		t.Errorf("Fail to return the correct string representation of the expression tree\n")
	}
}
