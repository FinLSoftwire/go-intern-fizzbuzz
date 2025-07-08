package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter an upper bound for FizzBuzz: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	upperBoundInput, err := strconv.Atoi(scanner.Text())
	if err == nil {
		fmt.Errorf("Expected an integer, instead got %s", scanner.Text())
	}
	fmt.Println(FizzBuzz(1, upperBoundInput))
}

func insertBeforeFirstB(wordsSlice []string, newWord string) (newWordsSlice []string) {
	for wordIndex, currentWord := range wordsSlice {
		if currentWord[0] == 'B' {
			newWordsSlice = append(append(newWordsSlice, newWord), wordsSlice[wordIndex:]...)
			return
		} else {
			newWordsSlice = append(newWordsSlice, currentWord)
		}
	}
	newWordsSlice = append(newWordsSlice, newWord)
	return
}

func FizzBuzz(minimumBound int, maximumBound int) (fbOutput string) {
	for currentNumber := minimumBound; currentNumber <= maximumBound; currentNumber++ {
		applicableWordsSlice := make([]string, 0)
		if currentNumber%3 == 0 {
			applicableWordsSlice = append(applicableWordsSlice, "Fizz")
		}
		if currentNumber%5 == 0 {
			applicableWordsSlice = append(applicableWordsSlice, "Buzz")
		}
		if currentNumber%7 == 0 {
			applicableWordsSlice = append(applicableWordsSlice, "Bang")
		}
		if currentNumber%11 == 0 {
			applicableWordsSlice = []string{"Bong"}
		}
		if currentNumber%13 == 0 {
			applicableWordsSlice = insertBeforeFirstB(applicableWordsSlice, "Fezz")
		}
		if currentNumber%17 == 0 {
			slices.Reverse(applicableWordsSlice)
		}
		if len(applicableWordsSlice) == 0 {
			fbOutput += strconv.Itoa(currentNumber) + "\n"
		} else {
			fbOutput += strings.Join(applicableWordsSlice, "") + "\n"
		}
	}
	return
}
