package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mergedTemps := make(map[int][]float64)

	for _, temp := range temperatures {
		var key int
		if temp >= 0 {
			key = int(math.Floor(temp/10) * 10)
		} else {
			key = int(math.Ceil(temp/10) * 10)
		}
		mergedTemps[key] = append(mergedTemps[key], temp)
	}

	fmt.Println(mergedTemps)
}
