package main

import (
	"fmt"
)

func main() {

	num := make([]int, 10)
	// num := []int{}

	for i, _ := range num {
		num[i] = i + 1
	}
	fmt.Println(num)

}

// while loop implementation in Go
// i := 0

// for i < 10 {
// 	fmt.Println(i)
// 	i++
// }

// For loop implementation in Go
// for i := 0; i < 10; i++ {
// 	fmt.Println(i)
// }
