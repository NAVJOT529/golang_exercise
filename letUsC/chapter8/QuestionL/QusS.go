package main

import (
	"fmt"
	"math"
)

/*
*This function is used to find the distance between last point from first point
 */
func main() {
	var x [10]int
	var y [10]int

	for i := 0; i < 10; i++ {
		fmt.Println("Enter the value of point x and y")
		fmt.Scanln(&x[i], &y[i])
	}
	distance := 0.0
	for i := 0; i < 9; i++ {
		distance = distance + math.Sqrt(math.Pow(float64(x[i+1]-x[i]), 2)+math.Pow(float64(y[i+1]-y[i]), 2))
	}
	fmt.Println("The distance of last point from the first point is", distance)
	return
}
