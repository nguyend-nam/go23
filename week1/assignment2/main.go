package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	utils "github.com/nguyend-nam/go23/week1/utils"
	slices "golang.org/x/exp/slices"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		if isTypeValid(args[1]) {
			if isInputsValid(args[1][1:], args[2:]) {
				inputs := args[2:]
				var numericInputs []string
				var alphabetInputs []string

				for i := 0; i < len(inputs); i++ {
					isNumeric := utils.IsNumeric(inputs[i])
					hasAlphabet := utils.HasAlphabet(inputs[i])

					if isNumeric {
						numericInputs = append(numericInputs, inputs[i])
					} else if hasAlphabet {
						alphabetInputs = append(alphabetInputs, inputs[i])
					}
				}

				sortedNumericStrings := sortNumericStrings(numericInputs)
				sortedAlphabetStrings := sortAlphabetStrings(alphabetInputs)

				concatArray := append(sortedNumericStrings, sortedAlphabetStrings...)
				for _, str := range concatArray {
					fmt.Printf(str + " ")
				}
			}
		}
	}
}

func isTypeValid(input string) bool {
	validTypes := []string{"number", "string", "mix"}

	if input[0:1] == "-" {
		if slices.Contains(validTypes, input[1:]) {
			return true
		} else {
			fmt.Println("Invalid type")
			return false
		}
	} else {
		fmt.Println("Invalid type")
		return false
	}
}

func isInputsValid(inputType string, inputs []string) bool {
	if inputType == "number" {
		for i := 0; i < len(inputs); i++ {
			if !utils.IsNumeric(inputs[i]) {
				fmt.Println("Invalid input")
				return false
			}
		}

		return true
	} else if inputType == "string" {
		for i := 0; i < len(inputs); i++ {
			if !utils.HasAlphabet(inputs[i]) {
				fmt.Println("Invalid input")
				return false
			}
		}

		return true
	}

	return true
}

func sortNumericStrings(array []string) []string {
	sortedArray := array

	sort.Slice(sortedArray, func(i, j int) bool {
		number1, _ := strconv.ParseFloat(sortedArray[i], 64)
		number2, _ := strconv.ParseFloat(sortedArray[j], 64)
		return number1 < number2
	})

	return sortedArray
}

func sortAlphabetStrings(array []string) []string {
	sortedArray := array

	sort.Strings(sortedArray)

	return sortedArray
}
