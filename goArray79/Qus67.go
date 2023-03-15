/* Write a golang program to find subArray which has the largest sum in a given circular array of integers. */
package main

import (
	"fmt"
	"math"
)

/*
*This function is used to find subArray which has largest sum in a given circularly array
 */
func main() {
	arr := []int{1, -2, 3, 0, 7, 8, 1, 2, -3}
	fmt.Println(arr)
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	currentMax := arr[0]
	max := arr[0]
	currentMin := arr[0]
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		currentMax = int(math.Max(float64(currentMax+arr[i]), float64(arr[i])))
		max = int(math.Max(float64(max), float64(currentMax)))
		currentMin = int(math.Min(float64(currentMin+arr[i]), float64(arr[i])))
		min = int(math.Min(float64(currentMin), float64(min)))
	}
	result := int(math.Max(float64(max), float64(sum-min)))
	fmt.Println(result)
	return
}

/*
https://www.geeksforgeeks.org/maximum-contiguous-circular-sum/ */
