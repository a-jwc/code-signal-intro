package main

import (
	"fmt"
)

func solution(picture []string) []string {
	yBorder := ""
	for i, v := range picture {
		picture[i] = "*" + v + "*"
	}
	pictureLength := len(picture[0])
	for i := 0; i < pictureLength; i++ {
		yBorder += "*"
	}
	picture = append([]string{yBorder}, picture...)
	return append(picture, yBorder)
	
}

func main() {
	picture := []string{"abc", "ded"}
	fmt.Println(solution(picture))
}
