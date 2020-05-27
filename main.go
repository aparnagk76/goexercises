package main

import (
	"fmt"

	arr "github.com/aparnagk76/goexercises/arrays"
)

func main() {
	a := []int{12, 45, 15, 25, 46, 50, 3, 102, 20}
	b := 9
	x, y := arr.FindTwoNumbers(a, b)
	fmt.Println(x, y)

}
