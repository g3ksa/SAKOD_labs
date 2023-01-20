package main

import (
	"fmt"
	"lab6/LinkedList"
	"math/rand"
	"time"
)

func calcMinAndMax(arr []LinkedList.LinkedList, N int) (int, int) {
	maxLen, minLen := 0, N
	for _, value := range arr {
		if value.Length > maxLen {
			maxLen = value.Length
		}
		if value.Length < minLen {
			minLen = value.Length
		}
	}
	return minLen, maxLen
}

func main() {
	rand.Seed(time.Now().UnixNano())
	N := 5000
	var array = make([]LinkedList.LinkedList, N/100)
	for i := 0; i < N; i++ {
		temp := rand.Intn(N-1) + 1
		array[temp%(N/100)].PushBack(temp)
	}
	// for _, list := range array {
	// 	list.Print()
	// }
	// fmt.Println()
	fmt.Println(calcMinAndMax(array, N))
}
