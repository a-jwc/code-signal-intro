package main

import (
	"fmt"
)

func solution(inputArray []int) int {
	max := getMax(inputArray)
	candidateArray := []int{}

	max = max + 1
	for i := 1; i <= max; i++ {
		if !doesValueExist(i, inputArray) {
			candidateArray = append(candidateArray, i)
		}
	}
	retVal := getMax(candidateArray)
	defaultValue := retVal
	found := false
	for _, i := range candidateArray {
		for _, v := range inputArray {
			if i > v {
				continue
			}
			if v > i && v % i != 0 && i != v {
				retVal = i
				found = true
			} else if v < i && i % v != 0 && i != v {
				retVal = i
				found = true
			} else {
				retVal = defaultValue
				found = false
				break
			}
		}
		if found {
			return retVal
		}
	}
	return retVal
}

func doesValueExist(value int, arr []int) bool {
	for _, v := range arr {
		if value == v {
			return true
		}
	}
	return false
}

func getMax(arr []int) (max int) {
	max = arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return
}

func main() {
	inputArray := []int{2, 3}
	fmt.Println(solution(inputArray))
}