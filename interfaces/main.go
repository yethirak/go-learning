package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// receiver value variable "eb" can be removed above since it's not being used in func
	// custom logic for generating English greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// receiver value variable "sb" can be removed above since it's not being used in func
	// custom logic for generating Spanish greeting
	return "Hola!"
}
