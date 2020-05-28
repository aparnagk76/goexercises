package main

import (
	"fmt"

	arr "github.com/aparnagk76/goexercises/arrays"
)

func main() {
	// a := []int{12, 45, 15, 25, 46, 50, 3, 102, 20}
	// b := 9
	// x, y := arr.FindTwoNumbers(a, b)
	// fmt.Println(x, y)
	// cars := []int{14, 12, 7, 6, 9, 5, 13}
	// Covered := 3
	// result := arr.FindMinRoofLength(cars, Covered)
	// fmt.Println(result)
	//c := []int{23, 8, 20, 6, 11, 23, 8, 20, 16, 4, 12, 23, 8, 18}
	//minTime, maxTime, avg := arr.GetSleepTime(c)
	//fmt.Println(minTime, maxTime, avg)
	a := []int{5, 9, 76, 35, 22, 85, 45, 60, 35}
	var b, c int = 30, 60
	heaterCount, coolerCount := arr.GetCount(a, b, c)
	fmt.Println(heaterCount, coolerCount)

}
