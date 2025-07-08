package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(FizzBuzz(1, 20))
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
		if len(applicableWordsSlice) == 0 {
			fbOutput += strconv.Itoa(currentNumber) + "\n"
		} else {
			fbOutput += strings.Join(applicableWordsSlice, "") + "\n"
		}
	}
	return
}
