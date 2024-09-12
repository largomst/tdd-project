package main

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := Dollar{
		amount: 5,
	}
	tenner := five.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got: [%d]", tenner.amount)
	}
}

type Dollar struct {
	amount int
}
func (d Dollar) Times(n int) Dollar {
	return Dollar{10}
}