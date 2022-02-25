package main

import(
	"fmt"
	"sort"
)

func solution(statues []int) int {
	sort.Ints(statues)
	count := 0
	
	for i := 0; i < len(statues) - 1; i++ {
		dif := statues[i + 1] - statues[i] - 1
		if dif > 0 {
			count += dif
		}
	}

	return count
}

func main() {
	arr := []int{6, 2, 3, 8, 10}
	fmt.Println(solution(arr))
}