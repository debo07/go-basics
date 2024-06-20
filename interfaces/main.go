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

	var b1 bot = englishBot{}
	fmt.Println("Greetings from b1: ",b1.getGreeting())
	printGreeting(b1)

	var b2 bot = spanishBot{}
	fmt.Println("Greetings fron b2: ", b2.getGreeting())
	printGreeting(b2)
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

