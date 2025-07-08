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

func TestBuzz(t *testing.T) {
	result := FizzBuzz(5, 5)
	if result != "Buzz\n" {
		t.Errorf("Expected Buzz, instead found: %s", result)
	}
}

func TestFizzAndBuzz(t *testing.T) {
	result := FizzBuzz(15, 15)
	if result != "FizzBuzz\n" {
		t.Errorf("Expected FizzBuzz, instead found: %s", result)
	}
}
