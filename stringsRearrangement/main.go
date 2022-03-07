package main

import (
	"fmt"
)

func solution(inputArray []string) bool {
	for perms := range GeneratePermutationsString(inputArray) {
		fmt.Println(perms)
		for i := 0; i < len(perms)-1; i++ {
			currWord := perms[i]
			nextWord := perms[i+1]

			diffCount := 0
			breakFlag := false
			for i := 0; i < len(currWord); i++ {
				if currWord[i] != nextWord[i] {
					diffCount += 1
				}
			}
			if diffCount != 1 {
				breakFlag = true
				break
			}
			if breakFlag {
				break
			} else if !breakFlag && i == len(perms)-2 {
				return true
			}
		}
	}
	return false
}

// source: https://github.com/Ramshackle-Jamathon/go-quickPerm
func GeneratePermutationsString(data []string) <-chan []string {
	c := make(chan []string)
	go func(c chan []string) {
		defer close(c)
		permutateString(c, data)
	}(c)
	return c
}

func permutateString(c chan []string, inputs []string) {
	output := make([]string, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]string, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func main() {
	inputArray := []string{"zzzzab", 
	"zzzzbb", 
	"zzzzaa"}
	fmt.Println(solution(inputArray))
}
