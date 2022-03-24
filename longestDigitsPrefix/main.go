package main

import (
	"fmt"
	"regexp"
)

func solution(inputString string) string {
	currDigits := ""
	regex, err := regexp.Compile(`\d`)

	if err != nil {
		return err.Error()
	}

	for i, v := range inputString {
		match := regex.MatchString(string(v))
		if !match && i == 0 {
			return ""
		}
		if match {
			currDigits += string(v)
		} else {
			return currDigits
		}
	}
	return currDigits
}

func main() {
	inputString := " 12sdf3df1s235s4aa123s4"
	fmt.Println(solution(inputString))
}
