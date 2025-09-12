package main

import "fmt"

func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			right = mid - 1
		}
		if arr[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print(binarySearch(arr, 10))
}
