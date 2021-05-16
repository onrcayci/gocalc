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
			t.Errorf("Fail to detect operand %s as an operand!\n", operands[i])
		}
	}
}

func TestGetOperatorPrecedence_Power(t *testing.T) {
	precedence := grammar.GetOperatorPrecedence("^")
	if precedence != 3 {
		t.Errorf("Fail to detect precedence of '^'\nExpected: 3\nOutcome: %d\n", precedence)
	}
}

func TestGetOperatorPrecedence_MultiplyDivide(t *testing.T) {
	operators := [2]string{"*", "/"}
	for i := 0; i < len(operators); i++ {
		precedence := grammar.GetOperatorPrecedence(operators[i])
		if precedence != 2 {
			t.Errorf("Fail to detect precedence of %s\nExpected: 2\nOutcome: %d\n", operators[i], precedence)
		}
	}
}

func TestGetOperatorPrecedence_PlusMinus(t *testing.T) {
	operators := [2]string{"+", "-"}
	for i := 0; i < len(operators); i++ {
		precedence := grammar.GetOperatorPrecedence(operators[i])
		if precedence != 1 {
			t.Errorf("Fail to detect precedence of %s\nExpected: 1\nOutcome: %d\n", operators[i], precedence)
		}
	}
}

func TestGetOperatorPrecedence_OperandsParantheses(t *testing.T) {
	characters := [3]string{"12", "(", ")"}
	for i := 0; i < len(characters); i++ {
		precedence := grammar.GetOperatorPrecedence(characters[i])
		if precedence != 0 {
			t.Errorf("Fail to detect precedence of %s\nExpected: 0\nOutcome: %d\n", characters[i], precedence)
		}
	}
}
