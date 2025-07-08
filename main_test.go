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

func TestBang(t *testing.T) {
	result := FizzBuzz(7, 7)
	if result != "Bang\n" {
		t.Errorf("Expected Bang, instead found: %s", result)
	}
}

func TestFizzAndBang(t *testing.T) {
	result := FizzBuzz(21, 21)
	if result != "FizzBang\n" {
		t.Errorf("Expected FizzBang, instead found: %s", result)
	}
}

func TestFizzBuzzBang(t *testing.T) {
	result := FizzBuzz(105, 105)
	if result != "FizzBuzzBang\n" {
		t.Errorf("Expected FizzBuzzBang, instead found: %s", result)
	}
}

func TestBong(t *testing.T) {
	result := FizzBuzz(11, 11)
	if result != "Bong\n" {
		t.Errorf("Expected Bong, instead found: %s", result)
	}
}

func TestBongOnly(t *testing.T) {
	result := FizzBuzz(33, 33)
	if result != "Bong\n" {
		t.Errorf("Expected Bong, instead found: %s", result)
	}
}
