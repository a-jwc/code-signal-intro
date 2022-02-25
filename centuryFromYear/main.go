package main

import (
	"fmt"
	"strconv"
)

func solution(year int) int {
	yearString := strconv.Itoa(year)
	yearLength := len(yearString)

	if yearLength <= 2 || yearString == "100" {
		return 1
	}

	lastChars := yearString[yearLength-2:]

	switch {
	case yearLength == 3:
		{
			firstChar := yearString[0:1]
			return convertToInt(firstChar, lastChars)
		}
	case yearLength == 4:
		{
			firstTwoChars := yearString[0:2]
			return convertToInt(firstTwoChars, lastChars)
		}
	}
	return 0
}

func convertToInt(first string, last string) int {
	if last == "00" {
		intVar, _ := strconv.Atoi(first)
		return intVar
	} else {
		intVar, _ := strconv.Atoi(first)
		return intVar + 1
	}
}

func main() {
	fmt.Println(solution(2001))
}
