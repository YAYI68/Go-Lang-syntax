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
	en := englishBot{}
	sp := spanishBot{}

	fmt.Println(printGreeting(en))
	fmt.Println(printGreeting(sp))
}

func printGreeting(b bot) string {
	return b.getGreeting()
}

func (en englishBot) getGreeting() string {
	return "Hello world"
}

func (sb spanishBot) getGreeting() string {
	return "Hoolla world !"
}
