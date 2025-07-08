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
	integerSliceOperationMapping map[int]func(currentWordsSlice *stringSlice)
	orderedRuleModulos           []int
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
	defaultRuleSet := initialiseDefaultRuleSet()
	fmt.Println(defaultRuleSet.FizzBuzz(1, upperBoundInput))
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

func (currentWordsSlice *stringSlice) setToSingleElement(word string) {
	*currentWordsSlice = stringSlice{word}
}

func initialiseDefaultRuleSet() RuleSet {
	defaultRuleSet := RuleSet{}
	defaultRuleSet.integerSliceOperationMapping = make(map[int]func(currentWordsSlice *stringSlice))
	defaultRuleSet.integerSliceOperationMapping[3] = func(currentWordsSlice *stringSlice) { currentWordsSlice.appendInPlace("Fizz") }
	defaultRuleSet.integerSliceOperationMapping[5] = func(currentWordsSlice *stringSlice) { currentWordsSlice.appendInPlace("Buzz") }
	defaultRuleSet.integerSliceOperationMapping[7] = func(currentWordsSlice *stringSlice) { currentWordsSlice.appendInPlace("Bang") }
	defaultRuleSet.integerSliceOperationMapping[11] = func(currentWordsSlice *stringSlice) { currentWordsSlice.setToSingleElement("Bong") }
	defaultRuleSet.integerSliceOperationMapping[13] = func(currentWordsSlice *stringSlice) { currentWordsSlice.insertBeforeFirstB("Fezz") }
	defaultRuleSet.integerSliceOperationMapping[17] = func(currentWordsSlice *stringSlice) { slices.Reverse(*currentWordsSlice) }
	defaultRuleSet.orderedRuleModulos = []int{3, 5, 7, 11, 13, 17}
	return defaultRuleSet
}

func (currentRuleSet *RuleSet) FizzBuzz(minimumBound int, maximumBound int) (fbOutput string) {
	for currentNumber := minimumBound; currentNumber <= maximumBound; currentNumber++ {
		applicableWordsSlice := stringSlice(make([]string, 0))
		for _, ruleInteger := range currentRuleSet.orderedRuleModulos {
			if currentNumber%ruleInteger == 0 {
				currentRuleSet.integerSliceOperationMapping[ruleInteger](&applicableWordsSlice)
			}
		}
		if len(applicableWordsSlice) == 0 {
			fbOutput += strconv.Itoa(currentNumber) + "\n"
		} else {
			fbOutput += strings.Join(applicableWordsSlice, "") + "\n"
		}
	}
	return
}
