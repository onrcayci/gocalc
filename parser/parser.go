package parser

import (
	"strings"
	"text/scanner"
)

func Tokenize(fileContents string) []string {

	// initialize the return array
	var tokens []string

	// initialize the scanner
	var scan scanner.Scanner
	scan.Init(strings.NewReader(fileContents))

	// loop over the input string to tokenize it
	for token := scan.Scan(); token != scanner.EOF; token = scan.Scan() {
		stringToken := scan.TokenText()
		tokens = append(tokens, stringToken)
	}

	// return the tokenized string array
	return tokens
}
