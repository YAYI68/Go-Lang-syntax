package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSome(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("episodes", &wg)

	wg.Wait()
	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "episodes") {
		t.Errorf("Expected to find episodes, got %s", output)
	}
}
