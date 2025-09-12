package main

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func sorting(arr []int, low, high int) []int {
	if low < high {
		p := partition(arr, low, high)
		sorting(arr, low, p-1)
		sorting(arr, p+1, high)
	}
	return arr
}

func quickSort(arr []int) []int {
	return sorting(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{2, 0, 7, 1, 5, 3, 6, 4}
	fmt.Println(quickSort(arr))
}
