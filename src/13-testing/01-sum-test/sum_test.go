package main

import "testing"

func TestSumAPlusZeroShouldReturnA(t *testing.T) {
	result := Sum(5, 0)
	expected := 5

	if result != expected {
		t.Errorf("Expected: %d, got: %d", expected, result)

	}
}

func TestSumAPlusBShouldReturnTheSum(t *testing.T) {
	result := Sum(5, 10)
	expected := 15

	if result != expected {
		t.Errorf("Expected: %d, got: %d", expected, result)

	}
}
