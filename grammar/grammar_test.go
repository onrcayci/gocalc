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

func TestIsOperator_ReturnsFalse(t *testing.T) {
	operands := [2]string{"1", "123.342"}
	for i := 0; i < len(operands); i++ {
		result := grammar.IsOperator(operands[i])
		if result {
			t.Errorf("Fail to detect operand %s as an operand!\n", operands[i]);
		}
	}
}

func TestGetOperatorPrecedence_Power(t *testing.T) {
	precendence := grammar.GetOperatorPrecedence("^")
	if precendence != 3 {
		t.Errorf("Fail to detect precedence of '^'\nExpected: 3\nOutcome: %T\n", precendence)
	}
}
