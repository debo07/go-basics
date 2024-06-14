package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct {}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	// In real, it will be very custom logic
	return "Hi, there!"
}

func (eb spanishBot) getGreeting() string {
	// In real, it will be very custom logic
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

