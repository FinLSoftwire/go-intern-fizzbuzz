package main

import (
	"testing"
)

func TestFizz(t *testing.T) {
	result := FizzBuzz(3, 3)
	if result != "Fizz\n" {
		t.Errorf("Expected Fizz, instead found: %s", result)
	}
}
