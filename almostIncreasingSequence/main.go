package main

import (
	"fmt"
	"math"
	// "sort"
)

func oldSolution(sequence []int) bool {
	for i := range sequence {
		// create a copy of the slice using append, use `make` to preallocate a slice of the right capacity
		clone := append(make([]int, 0, len(sequence)), sequence...)
		temp := append(clone[:i], clone[i+1:]...)
		// * could not use SliceIsSorted because it returns true for equal values
		// if sort.SliceIsSorted(temp, func(i, j int) bool { return temp[i] < temp[j] }) {
		if isIncreasing(temp) {
			fmt.Println(temp)
			return true
		}
	}
	return false
}

// * source: https://stackoverflow.com/questions/43017251/solve-almostincreasingsequence-codefights
func solution(sequence []int) bool {
	max := int(math.Inf(-1))
	prevMax := int(math.Inf(-1))
	removed := false

	for i := 0; i < len(sequence); i++ {
		// * if first element in slice or if current element is greater than the max
		// * set the prevMax to max, set the max to the current element
		if max == int(math.Inf(-1)) || sequence[i] > max {
			fmt.Println("prevMax", prevMax)
			prevMax = max
			max = sequence[i]
		} else if prevMax == int(math.Inf(-1)) || sequence[i] > prevMax {
			// * if second element in the slice or current element is greater than prevMax
			// * if an element has already been removed, return false
			// * remove the current element, set max to current element
			if removed {
				return false
			}
			removed = true
			max = sequence[i]
		} else {
			// * if the current element is less than both prevMax and max, remove it and set removed to true
			if removed {
				return false
			}
			removed = true
		}
	}
	return true
}

func isIncreasing(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if j := i + 1; arr[j] <= arr[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 5, 2}
	fmt.Println(solution(arr))
}
