package main

import (
	"fmt"
)

func solution(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	yourStrongest := yourLeft
	yourWeakest := yourRight
	if yourRight > yourLeft {
		yourStrongest = yourRight
		yourWeakest = yourLeft
	}

	friendsStrongest := friendsLeft
	friendsWeakest := friendsRight
	if friendsRight > friendsLeft {
		friendsStrongest = friendsRight
		friendsWeakest = friendsLeft
	} 

	strongestBool := yourStrongest == friendsStrongest
	weakestBool := yourWeakest == friendsWeakest

	if strongestBool && weakestBool {
		return true
	}
	return false
}

func main() {
	yourLeft := 10
	yourRight := 15
	friendsLeft := 15
	friendsRight := 9

	fmt.Println(solution(yourLeft, yourRight, friendsLeft, friendsRight))
}
