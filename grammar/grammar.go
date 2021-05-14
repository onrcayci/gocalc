package grammar

func IsOperator(operator string) bool {
	switch operator {
	case "+", "-", "*", "/", "^", "(", ")":
		return true
	default:
		return false
	}
}

func GetOperatorPrecedence(operator string) int {
	switch operator {
	case "^":
		return 3
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	default:
		return 0
	}
}
