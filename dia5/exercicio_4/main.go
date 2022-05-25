package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SelSort(list []int) {
	timeInit := time.Now()
	for j := 0; j < len(list); j++ {
		minIndex := j
		for i := j; i < len(list); i++ {
			if list[i] < list[minIndex] {
				minIndex = i
			}
		}
		if list[j] > list[minIndex] {
			list[j], list[minIndex] = list[minIndex], list[j]
		}
	}
	fmt.Println("Selection time - ", time.Since(timeInit), "Number of elements:", len(list))
}

func InsSort(list []int) {
	timeInit := time.Now()
	for i := 1; i < len(list); i++ {
		testNumber := list[i]

		previousNumber := i - 1

		for previousNumber >= 0 && list[previousNumber] > testNumber {
			list[previousNumber+1] = list[previousNumber]
		}
		list[previousNumber+1] = testNumber
	}
	fmt.Println("Insertion Sort - ", time.Since(timeInit), "Number of elements:", len(list))
}

func BubSort(list []int) {
	timeInit := time.Now()

	for i := 0; i < len(list); i++ {
		for index := 0; index < len(list)-1; index++ {
			if list[index] > list[index+1] {
				list[index], list[index+1] = list[index+1], list[index]
			}
		}
	}
	fmt.Println("Bubble Sort - ", time.Since(timeInit), "Number of elements:", len(list))
}

func main() {
	variavel := rand.Perm(100)
	variavel2 := rand.Perm(1000)
	variavel3 := rand.Perm(10000)
	SelSort(variavel)
	InsSort(variavel)
	BubSort(variavel)
	fmt.Println()
	SelSort(variavel2)
	InsSort(variavel2)
	BubSort(variavel2)
	fmt.Println()
	SelSort(variavel3)
	InsSort(variavel3)
	BubSort(variavel3)
}
