package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type RuleSet struct {
	integerSliceOperationMapping map[int]func(newWord string)
}

type stringSlice []string

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

func (currentWordsSlice *stringSlice) insertBeforeFirstB(newWord string) {
	newWordsSlice := make([]string, 0)
	for wordIndex, currentWord := range *currentWordsSlice {
		if currentWord[0] == 'B' {
			*currentWordsSlice = append(append(newWordsSlice, newWord), (*currentWordsSlice)[wordIndex:]...)
			return
		} else {
			newWordsSlice = append(newWordsSlice, currentWord)
		}
	}
	*currentWordsSlice = append(newWordsSlice, newWord)
}

func (currentWordsSlice *stringSlice) appendInPlace(newWord string) {
	*currentWordsSlice = append(*currentWordsSlice, newWord)
}

func FizzBuzz(minimumBound int, maximumBound int) (fbOutput string) {
	for currentNumber := minimumBound; currentNumber <= maximumBound; currentNumber++ {
		applicableWordsSlice := stringSlice(make([]string, 0))
		if currentNumber%3 == 0 {
			applicableWordsSlice.appendInPlace("Fizz")
		}
		if currentNumber%5 == 0 {
			applicableWordsSlice.appendInPlace("Buzz")
		}
		if currentNumber%7 == 0 {
			applicableWordsSlice.appendInPlace("Bang")
		}
		if currentNumber%11 == 0 {
			applicableWordsSlice = []string{"Bong"}
		}
		if currentNumber%13 == 0 {
			applicableWordsSlice.insertBeforeFirstB("Fezz")
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
