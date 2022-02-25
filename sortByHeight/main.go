package main

import (
	"fmt"
	"sort"
)

type Items struct {
	list   []int
	length int
	people bool
}

func solution(a []int) []int {
	countMap := make(map[int]Items)
	peopleList := []int{}
	retList := []int{}

	countMapIndex := 0
	listCount := 0
	peopleBool := false
	prevItem := 0
	tempList := []int{}

	for _, v := range a {
		if (prevItem == -1 && v != -1) || (prevItem != -1 && v == -1) {
			tempList = nil
			countMapIndex += 1
			listCount = 0
		}
		tempList = append(tempList, v)
		if v == -1 {
			peopleBool = false
			listCount += 1
		} else {
			peopleBool = true
			listCount += 1
		}
		countMap[countMapIndex] = Items{tempList, listCount, peopleBool}
		prevItem = v
	}

	// countMap[0] = Items{[]int{-1}, 1, false}
	// countMap[1] = Items{[]int{150, 190, 170}, 3, true}
	// countMap[2] = Items{[]int{-1, -1}, 2, false}
	// countMap[3] = Items{[]int{160, 180}, 2, true}

	for _, v := range countMap {
		if v.people {
			peopleList = append(peopleList, v.list...)
		}
	}
	sort.Ints(peopleList)
	// fmt.Println(peopleList)

	j := 0
	endCount := 0
	for i := 0; i <= len(countMap); i++ {
		if countMap[i].people {
			endCount += countMap[i].length
			retList = append(retList, peopleList[j:endCount]...)
			j += countMap[i].length
		} else {
			retList = append(retList, countMap[i].list...)
		}
	}

	// ! weird behavior; ordering is nondeterministic
	// for i := range countMap {
	// 	if countMap[i].people {
	// 		endCount += countMap[i].length
	// 		retList = append(retList, peopleList[j:endCount]...)
	// 		j += countMap[i].length
	// 	} else {
	// 		retList = append(retList, countMap[i].list...)
	// 	}
	// }

	return retList
}

func main() {
	test1 := []int{-1, 150, 190, 170, -1, -1, 160, 180}
	// test6 := []int{23, 54, -1, 43, 1, -1, -1, 77, -1, -1, -1, 3}
	fmt.Println(solution(test1))
}
