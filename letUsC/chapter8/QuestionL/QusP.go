package main

import (
	"fmt"
	"math"
)

/*
*This function is used to find the largest land in the list
 */
func main() {
	var a [6]float64
	var b [6]float64
	var angle [6]float64
	var area [6]float64
	fmt.Println("Enter the a side b side and angle of the plot")
	for i := 0; i < 6; i++ {
		fmt.Scanln(&a[i], &b[i], &angle[i])
		area[i] = (1 / 2) * a[i] * b[i] * math.Sin(angle[i])
	}
	max := area[0]
	num := 1
	for i := 1; i < 6; i++ {
		if area[i] > max {
			max = area[i]
			num = i + 1
		}
	}
	fmt.Println("The largest plot in the list is plot number", num, "with the area of ", max)
	return
}
