package parser

import (
	"fmt"
	"strings"
	"text/scanner"
)

func ParseContent(fileContents string) {

	// initialize the scanner
	var scan scanner.Scanner
	scan.Init(strings.NewReader(fileContents))

	i := 0

	for token := scan.Scan(); token != scanner.EOF; token = scan.Scan() {
		fmt.Printf("token %d: %s \n", i, scan.TokenText())
		i++
	}
}
