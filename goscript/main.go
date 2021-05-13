package main

import (
	"fmt"
	"os"
)


func main() {

	// get the command line arguments
	args := os.Args[1:]
	fmt.Println(args)
}