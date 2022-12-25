package main

import "fmt"

func main() {
	printSomething("This is my first message!")
	printSomething("This is my second message!")
}

func printSomething(s string) {
	fmt.Println(s)
}
