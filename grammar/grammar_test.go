package grammar_test

import (
	"testing"

	"github.com/onrcayci/gocalculator/grammar"
)

func TestIsOperator_ReturnsTrue(t *testing.T) {
	operators := [7]string{"+", "-", "*", "/", "^", "(", ")"}
	for i := 0; i < len(operators); i++ {
		result := grammar.IsOperator(operators[i])
		if !result {
			t.Errorf("Fail to detect the operator %s as an operator!\n", operators[i])
		}
	}
}
