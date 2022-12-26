package main

import "testing"

func Test_dine(t *testing.T) {

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("expected 5, got %v", len(orderFinished))
		}

	}

}
