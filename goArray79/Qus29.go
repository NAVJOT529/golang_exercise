/* 29. Write a golang program to compute the average value of an array of integers except the largest and smallest values. */
package main

import "fmt"

/*
*This function is used to find the average values of the elements of an array expect smallest and largest values
 */
func main() {
	var arr = []int{55, 44, 80, 90, 100, 25, 24, 23, 45}
	fmt.Println(arr)
	max := arr[0]
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	sum := 0
	for i := range arr {
		sum = sum + arr[i]
	}
	sum = sum - max - min
	average := sum / (len(arr) - 2)
	fmt.Println("average value of the element of the array expect largest and smallest value is", average)
	return
}
