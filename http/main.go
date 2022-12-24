package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// os.WriteFile("index.html", bs, 0666)
	// fmt.Println(string(bs))
}

type logWriter struct{}

// Custom Interface Using Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("length of bytes written", len(bs))
	return len(bs), nil
}
