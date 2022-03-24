package main

import (
	"fmt"
	"regexp"
)

func solution(inputString string) string {
	regex, err := regexp.Compile(`\d`)
	if err != nil {
		return err.Error()
	}
	for _, v := range inputString {
		match := regex.MatchString(string(v))
		if match {
			return string(v)
		}
	}
	return ""
}

func main() {
	inputString := "var_1__Int"
	fmt.Println(solution(inputString))
}