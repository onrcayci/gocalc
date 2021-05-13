package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/onrcayci/goscript/parser"
)

func main() {

	// get the command line arguments
	filename := os.Args[1]

	// open the file for reading
	file, err := os.Open(filename)

	// check for errors
	if err != nil {
		panic(err)
	}

	// close file with defer
	defer file.Close()

	// initialize scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// initialize the array to store the lines
	var lines string

	// loop over the file and scan each line
	for scanner.Scan() {
		lines += scanner.Text()
	}

	// check for read errors
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", lines)

	// parse the contents of the file
	parser.ParseContent(lines)
}
