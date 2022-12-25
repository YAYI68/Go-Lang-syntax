package main

import (
	"fmt"
	"sync"
)

func main() {
	// Creating my variable for waitGroup
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"episilon",
	}

	//    Add number of words to wait for
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("This is my second message!", &wg)
}

// Add a wg pointer as parameter to print something
// Dercrement wg by calling  wg.Done()
func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
