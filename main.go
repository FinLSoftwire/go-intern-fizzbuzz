package main

import (
	"fmt"
	"strconv"
	"strings"
)

// This is our main function, this executes by default when we run the main package.
func main() {
	fmt.Println(FizzBuzz(1, 10))
}

func FizzBuzz(minimumBound int, maximumBound int) (fbOutput string) {
	for currentNumber := minimumBound; currentNumber <= maximumBound; currentNumber++ {
		applicableWordSlice := make([]string, 0)
		if currentNumber%3 == 0 {
			applicableWordSlice = append(applicableWordSlice, "Fizz")
		}
		if len(applicableWordSlice) == 0 {
			fbOutput += strconv.Itoa(currentNumber) + "\n"
		} else {
			fbOutput += strings.Join(applicableWordSlice, "") + "\n"
		}
	}
	return
}
