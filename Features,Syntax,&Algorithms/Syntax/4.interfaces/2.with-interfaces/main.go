package main

import (
	"fmt"
)

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

func (englishBot) getGreeting() string {
	return "this is english bot => Hello!"
}

func (spanishBot) getGreeting() string {
	return "esto es bot español => Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
