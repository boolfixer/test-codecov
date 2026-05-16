package main

import "testing"

func TestCalculatorSum(t *testing.T) {
	c := &Calculator{}

	if got := c.Sum(2, 3); got != 5 {
		t.Fatalf("Sum(2, 3) = %d, want 5", got)
	}
}

func TestCalculatorSub(t *testing.T) {
	c := &Calculator{}

	if got := c.Sub(7, 4); got != 3 {
		t.Fatalf("Sub(7, 4) = %d, want 3", got)
	}
}

//func TestCalculatorMul(t *testing.T) {
//	c := &Calculator{}
//
//	if got := c.Mul(6, 5); got != 30 {
//		t.Fatalf("Mul(6, 5) = %d, want 30", got)
//	}
//}
