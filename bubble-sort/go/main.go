package main

import "fmt"

func main() {
	arr := []int{34, 7, 23, 32, 5, 62, 3, 9, 1, 18}
	fmt.Println("Unsorted array:", arr)
	fmt.Println("Sorted array:", bubbleSort(arr, false))
	fmt.Println("Descending order sorted array:", bubbleSort(arr, true))
}

func bubbleSort(array []int, desc bool) []int {
	sortOrder := 1
	if desc {
		sortOrder = -1
	}

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if (array[j]-array[j+1])*sortOrder > 0 {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
