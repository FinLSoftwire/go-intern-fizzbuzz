package main

import (
	"testing"
)

var defaultRuleSet, onlyFizzRuleSet RuleSet

func TestFizz(t *testing.T) {
	defaultRuleSet = initialiseDefaultRuleSet()
	onlyFizzRuleSet = initialiseSpecialisedRuleSet([]int{3})
	result := defaultRuleSet.FizzBuzz(3, 3)
	if result != "Fizz\n" {
		t.Errorf("Expected Fizz, instead found: %s", result)
	}
}

func TestBuzz(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(5, 5)
	if result != "Buzz\n" {
		t.Errorf("Expected Buzz, instead found: %s", result)
	}
}

func TestFizzAndBuzz(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(15, 15)
	if result != "FizzBuzz\n" {
		t.Errorf("Expected FizzBuzz, instead found: %s", result)
	}
}

func TestBang(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(7, 7)
	if result != "Bang\n" {
		t.Errorf("Expected Bang, instead found: %s", result)
	}
}

func TestFizzAndBang(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(21, 21)
	if result != "FizzBang\n" {
		t.Errorf("Expected FizzBang, instead found: %s", result)
	}
}

func TestFizzBuzzBang(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(105, 105)
	if result != "FizzBuzzBang\n" {
		t.Errorf("Expected FizzBuzzBang, instead found: %s", result)
	}
}

func TestBong(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(11, 11)
	if result != "Bong\n" {
		t.Errorf("Expected Bong, instead found: %s", result)
	}
}

func TestBongOnly(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(33, 33)
	if result != "Bong\n" {
		t.Errorf("Expected Bong, instead found: %s", result)
	}
}

func TestFezz(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(13, 13)
	if result != "Fezz\n" {
		t.Errorf("Expected Fezz, instead found: %s", result)
	}
}

func TestFezzBong(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(143, 143)
	if result != "FezzBong\n" {
		t.Errorf("Expected FezzBong, instead found: %s", result)
	}
}

func TestFezzBuzz(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(65, 65)
	if result != "FezzBuzz\n" {
		t.Errorf("Expected FezzBuzz, instead found: %s", result)
	}
}

func TestFizzFezzBuzz(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(195, 195)
	if result != "FizzFezzBuzz\n" {
		t.Errorf("Expected FizzFezzBuzz, instead found: %s", result)
	}
}

func TestFizzBuzzReverse(t *testing.T) {
	result := defaultRuleSet.FizzBuzz(255, 255)
	if result != "BuzzFizz\n" {
		t.Errorf("Expected BuzzFizz, instead found: %s", result)
	}
}

func TestLimitedRuleSet(t *testing.T) {
	result := onlyFizzRuleSet.FizzBuzz(3, 3)
	if result != "Fizz\n" {
		t.Errorf("Expected Fizz, instead found: %s", result)
	}
	result2 := onlyFizzRuleSet.FizzBuzz(5, 5)
	if result2 != "5\n" {
		t.Errorf("Expected 5, instead found: %s", result)
	}
}
