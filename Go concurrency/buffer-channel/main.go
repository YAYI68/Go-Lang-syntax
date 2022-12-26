package main

import (
	"fmt"
	"time"
)

func listenBuffer(ch chan int) {
	for {
		i := <-ch
		fmt.Printf("%v is printed ", i)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	// Addition of 10 create the buffer
	ch := make(chan int, 10)

	go listenBuffer(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel...")

	}

	fmt.Println("Done")
	close(ch)

}
