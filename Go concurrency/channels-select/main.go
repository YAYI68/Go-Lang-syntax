package main

import (
	"fmt"
	"time"
)

func server1(channel chan string) {

	for {
		time.Sleep(6 * time.Second)
		channel <- "This is from the server 1"
	}
}

func server2(channel chan string) {
	for {
		time.Sleep(3 * time.Second)
		channel <- "This is from the server 2"
	}
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Println("case 1", s1)
		case s2 := <-channel1:
			fmt.Println("case 2", s2)
		case s3 := <-channel2:
			fmt.Println("case 3", s3)
		case s4 := <-channel2:
			fmt.Println("case 4", s4)
		}

	}

}
