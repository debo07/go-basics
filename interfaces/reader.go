package main

import (
	"fmt"
	"io"
	"os"
)



func main() {
	args := os.Args
	fmt.Println(args[1:])

	if (len(args) < 2 || args[1] == "") {
		fmt.Println("Filename command line argument is missing. e.g. go run reader.go dummy.txt")
		os.Exit(1)
	}

	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error in opening file", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}