package main

import (
	"fmt"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func recSolution(inputString string) string {
	retString := ""
	revString := ""
	count := 0

	for i := 0; i < len(inputString); i++ {
		char := string(inputString[i])
		if char == "(" {
			count += 1
			revString = solution(inputString[i+1:])
			retString = retString + revString
			// i += len(revString) + count
			i = len(retString) + count * 2
		}
		if char != "(" && char != ")" {
			retString = retString + char
		}
		if char == ")" {
			retString = reverse(retString)
			return retString
		}
	}

	return retString
}

// source: https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/discuss/383670/JavaC%2B%2BPython-Tenet-O(N)-Solution
func solution(inputString string) string {
	strLen := len(inputString)
	stack := []int{}
	pair := make([]int, strLen)

	for i := 0; i < strLen; i++ {
		char := string(inputString[i])
		if char == "(" {
			stack = append(stack, i)
		}
		if char == ")" {
			n := len(stack) - 1
			j := stack[n:]
			stack = stack[:n]
			pair[i] = j[0]
			pair[j[0]] = i
		}
	}

	result := ""
	d := 1
	for i := 0; i < strLen; i += d {
		char := string(inputString[i])
		if char == "(" || char == ")" {
			i = pair[i]
			d = -d
		} else {
			result += char
		}
	}
	return result
}

func main() {
	// fmt.Println(solution("foo(bar)baz(blim)"))
	fmt.Println(solution("foo(bar(baz))blim")) // foo(bar zab das)blim => foo sad baz rab blim
	fmt.Println(solution("foo(bar(baz)(sad))blim")) // foo(bar zab das)blim => foo sad baz rab blim
}
