package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 3

	copy(arr[i:], arr[i+1:])
	// arr[:len(arr)-1] = nil - в случае, если в срезе указатели
	arr = arr[:len(arr)-1]

	fmt.Println(arr)
}
