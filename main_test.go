package main

import "testing"

func TestCalculatorSum(t *testing.T) {
	c := &Calculator{}

	if got := c.Sum(2, 3); got != 5 {
		t.Fatalf("Sum(2, 3) = %d, want 5", got)
	}
}
