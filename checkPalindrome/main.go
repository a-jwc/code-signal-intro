package main

import (
	"fmt"
)

func solution(inputString string) bool {
	inputStringLen := len(inputString)
	// return true if inputString is len 1
	if inputStringLen == 1 {
		return true
	}

	fmt.Println(inputStringLen / 2)
	if inputStringLen % 2 == 0 {
		j := inputStringLen / 2
		i := j - 1
		return check(i, j, inputString)
	} else {
		middle := inputStringLen / 2
		i := middle - 1
		j := middle + 1
		return check(i, j, inputString)
	}
}

func check(i int, j int, inputString string) bool {
	for i >= 0 {
		fmt.Printf("i: %c j: %c \n", inputString[i], inputString[j])
		if inputString[i] != inputString[j] {
			return false
		}
		i--
		j++
	}
	return true
}

func main() {
	fmt.Println(solution("aabadaa"))
}