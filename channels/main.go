package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	links := []string{
		"https://google.com/",
		"https://facebook.com/",
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://golang.org/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link,c)
	}

	for {
		  go checkLink(<- c,c) 
	}
      
	
	for l := range c {
        go func (link string) {
			time.Sleep(5 * time.Second)
			checkLink(link,c)
		} (l)
	}

	// for i := 0; i < len(links); i++ {
	//      fmt.Println(<- c)
	// }

}

func checkLink(link string,c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		os.Exit(1)
		return
	}
	fmt.Println(link + "  is success")
	c <- link

}
