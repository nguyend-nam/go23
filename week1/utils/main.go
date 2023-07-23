package utils

import "regexp"

func IsNumeric(input string) bool {
	return regexp.MustCompile(`^-?\d+(\.\d+)?$`).MatchString(input)
}

func HasAlphabet(input string) bool {
	return regexp.MustCompile(`[A-Za-z]`).MatchString(input)
}
