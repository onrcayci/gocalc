package utils

import (
	"bufio"
	"fmt"
	"os"
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

func ExecuteInput(tokens []string) {
	switch tokens[0] {
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Unknown Input!")
	}
}
