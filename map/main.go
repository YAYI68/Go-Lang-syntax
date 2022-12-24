package main

import (
	"fmt"
)

func main() {

	colors := make(map[string]string)

	colors["white"] = "ffff"
	colors["black"] = "000000"

	printColor(colors)
}

func printColor(c map[string]string) {
	for k, v := range c {
		fmt.Println("color " + k + " is " + v)
	}

}

// one way to create an empty map
// var colors map[string]string

// 2nd
// colors := map[string]string{
// 	"white": "ffff",
// 	"black": "000000",
// }

// delete(colors, "white")
