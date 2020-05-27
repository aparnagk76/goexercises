package arrays

import (
	"math"
	"sort"
)

func FindMinRoofLength(cars []int, k int) int {
	sort.Ints(cars)
	minRoofLength := math.MaxInt32
	currentRoofLength := 0
	for i := 0; i < len(cars)-k+1; i++ {
		currentRoofLength = cars[i-1+k] - cars[i] + 1
		if currentRoofLength < minRoofLength {
			minRoofLength = currentRoofLength
		}
	}
	return minRoofLength
}

