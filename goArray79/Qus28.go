/* 28. Write a golang program to get the difference between the largest and smallest values in an array of integers. The length of the array must be 1 and above. */
package main

import "fmt"

/*
*This function is used to find the difference between largest and smallest values in the array
 */
func main() {
	var arr = []int{2, 5, 1, 6, 9, 11}
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
	fmt.Println(min)
	fmt.Println(max)
	fmt.Println(max - min)
	return
}
