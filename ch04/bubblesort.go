// CH04 - Bubble Sort
package main

import (
	"fmt"
	"math/rand"
)

const max = 1000

func main() {
	fmt.Println("A bubble sort algorithm")
	list := generateRandArray()

	fmt.Println("List before sorting:", *list)
	bubbleSort(list)
}

func generateRandArray() *[max]int {

	rand.Seed(123010597100587)
	var arr [max]int

	for i := 0; i < max; i++ {
		arr[i] = rand.Intn(max)
	}

	return &arr
}

func bubbleSort(arr *[max]int) {

	unsortedUntilIndex := len(arr) - 1
	sorted := false

	list := *arr

	for sorted != true {
		sorted = true
		for i := 0; i < unsortedUntilIndex; i++ {
			if list[i] > list[i+1] {
				sorted = false
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
		unsortedUntilIndex -= 1
	}
	fmt.Println("Sorted array:", list)

}
