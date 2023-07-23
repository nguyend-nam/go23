package main

import (
	"fmt"
	"os"

	slices "golang.org/x/exp/slices"
)

func main() {
	args := os.Args

	if len(args) > 1 {
		fullname := constructFullname(args[1:])
		fmt.Println(fullname)
	}
}

func isCountyCodeValid(countryCode string) bool {
	validCountryCodes := []string{"US", "VN"}
	return slices.Contains(validCountryCodes, countryCode)
}

func isArgValid(arr []string) bool {
	if len(arr) == 3 || len(arr) == 4 {
		if (len(arr) == 3 && isCountyCodeValid(arr[2])) || len(arr) == 4 && isCountyCodeValid(arr[3]) {
			return true
		}
	}

	return false
}

func constructFullname(arr []string) string {
	var fullname string

	if isArgValid(arr) {
		if arr[len(arr)-1] == "US" {
			for i := 0; i < len(arr)-1; i++ {
				fullname += arr[i]
				if i < len(arr)-2 {
					fullname += " "
				}
			}
		}
		if arr[len(arr)-1] == "VN" {
			for i := len(arr) - 2; i >= 0; i-- {
				fullname += arr[i]
				if i > 0 {
					fullname += " "
				}
			}
		}
	} else {
		fmt.Println("Invalid args")
	}

	return fullname
}
