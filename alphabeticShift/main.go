package main

import (
	"fmt"
)

func solution(inputString string) string {
	retString := ""
	for _, v := range inputString {
		if int(v) == 122 {
			retString = retString + "a"
		} else {
			retString = retString + string(rune(int(v)+1))
		}
	}
	return retString
}

func main() {
	inputString := "abc"
	fmt.Println(solution(inputString))
}
