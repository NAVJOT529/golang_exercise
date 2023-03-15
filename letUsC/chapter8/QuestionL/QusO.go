package main

import (
	"fmt"
	"math"
)

/*
*This function is used to find the standard deviation and the mean
 */
func main() {
	var num = []int{-6, -12, 8, 13, 11, 6, 7, 2, -6, -9, -10, 11, 10, 9, 2}
	sum := 0
	for i := range num {
		sum = sum + num[i]
	}
	mean := sum / len(num)
	fmt.Println("mean is :", mean)
	add := 0
	for i := range num {
		add = add + int(math.Pow(float64(num[i]-mean), 2))
	}
	StandardDeviation := math.Sqrt(float64(add / len(num)))
	fmt.Println("Standard deviation is :", StandardDeviation)
	return
}
