package main

import (
	"fmt"
	"sync"
)

var msgList []string

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	var wg sync.WaitGroup

	msgList = append(msgList, "Hello, world!", "Hello, universe!", "Hello, cosmos!")

	wg.Add(len(msgList))
	for _, v := range msgList {
		// fmt.Println("msg", v)
		go updateMessage(v, &wg)
		go printMessage()
	}
	wg.Wait()
}
