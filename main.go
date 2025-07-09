package main

import (
	"bufio"
	. "fizzbuzz/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type RuleSet struct {
	integerSliceOperationMapping map[int]func(currentWordsSlice *StringSlice)
	orderedModuloList            []int
}

func main() {
	upperBoundInput := getFizzBuzzUpperBoundInput()
	var ruleSetInUse RuleSet
	integerArguments := getIntsFromArguments(os.Args)
	ruleSetInUse = initialiseSpecialisedRuleSet(integerArguments)
	fmt.Println(ruleSetInUse.FizzBuzz(1, upperBoundInput))
}

func getFizzBuzzUpperBoundInput() int {
	fmt.Println("Enter an upper bound for FizzBuzz: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if len(scanner.Text()) == 0 {
		return 100
	}
	return handleASCIIToIntConversion(scanner.Text())
}

func handleASCIIToIntConversion(asciiInput string) (integerEquivalence int) {
	integerEquivalence, err := strconv.Atoi(asciiInput)
	if err == nil {
		fmt.Errorf("Expected an integer, instead got %s", asciiInput)
	}
	return
}

func getIntsFromArguments(arguments []string) (returnIntegers []int) {
	if len(arguments) < 2 {
		returnIntegers = []int{3, 5, 7, 11, 13, 17}
		return
	}
	returnIntegers = make([]int, 0)
	for argumentIndex, currentArgumentString := range arguments {
		if argumentIndex < 1 {
			continue
		}
		returnIntegers = append(returnIntegers, handleASCIIToIntConversion(currentArgumentString))
	}
	return
}

func initialiseDefaultRuleSet() RuleSet {
	defaultRuleSet := RuleSet{}
	defaultRuleSet.integerSliceOperationMapping = make(map[int]func(currentWordsSlice *StringSlice))
	defaultRuleSet.integerSliceOperationMapping[3] = func(currentWordsSlice *StringSlice) { currentWordsSlice.AppendInPlace("Fizz") }
	defaultRuleSet.integerSliceOperationMapping[5] = func(currentWordsSlice *StringSlice) { currentWordsSlice.AppendInPlace("Buzz") }
	defaultRuleSet.integerSliceOperationMapping[7] = func(currentWordsSlice *StringSlice) { currentWordsSlice.AppendInPlace("Bang") }
	defaultRuleSet.integerSliceOperationMapping[11] = func(currentWordsSlice *StringSlice) { currentWordsSlice.SetToSingleElement("Bong") }
	defaultRuleSet.integerSliceOperationMapping[13] = func(currentWordsSlice *StringSlice) { currentWordsSlice.InsertBeforeFirstB("Fezz") }
	defaultRuleSet.integerSliceOperationMapping[17] = func(currentWordsSlice *StringSlice) { slices.Reverse(*currentWordsSlice) }
	defaultRuleSet.orderedModuloList = []int{3, 5, 7, 11, 13, 17}
	return defaultRuleSet
}

func initialiseSpecialisedRuleSet(desiredBuiltinRules []int) RuleSet {
	currentRuleSet := RuleSet{}
	currentRuleSet.integerSliceOperationMapping = make(map[int]func(currentWordsSlice *StringSlice))
	currentRuleSet.orderedModuloList = make([]int, 0)
	defaultRuleSet := initialiseDefaultRuleSet()
	// Rules should be applied in the order in ascending integer order
	slices.Sort(desiredBuiltinRules)
	for _, desiredRuleInteger := range desiredBuiltinRules {
		if slices.Contains(defaultRuleSet.orderedModuloList, desiredRuleInteger) {
			currentRuleSet.integerSliceOperationMapping[desiredRuleInteger] = defaultRuleSet.integerSliceOperationMapping[desiredRuleInteger]
			currentRuleSet.orderedModuloList = append(currentRuleSet.orderedModuloList, desiredRuleInteger)
		}
	}
	return currentRuleSet
}

func (currentRuleSet *RuleSet) FizzBuzz(minimumBound int, maximumBound int) (fbOutput string) {
	for currentNumber := minimumBound; currentNumber <= maximumBound; currentNumber++ {
		applicableWordsSlice := StringSlice(make([]string, 0))
		for _, ruleInteger := range currentRuleSet.orderedModuloList {
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
