package main

import "fmt"
var c, python bool 
var java string

func main() {
	name:= "bill"

	namePointer := &name

	fmt.Println(namePointer)
	printPointer(namePointer);
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer);
}