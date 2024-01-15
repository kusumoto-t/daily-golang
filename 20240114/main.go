package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	sortedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	target := 5

	index := BinarySearch(sortedSlice, target)

	if index != -1 {
		fmt.Printf("%d はスライスの %d番目にあります。\n", target, index)
	} else {
		fmt.Printf("%d はスライスにありませんでした。\n", target)
	}
}
