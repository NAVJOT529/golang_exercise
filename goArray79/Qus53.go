/* 53. Write a golang program to replace every element with the next greatest element (from right side) in a given array of integers. There is no element next to the last element, therefore replace it with -1. */
package main

import "fmt"

/*
*This func is used to to replace every element with the next greatest element in the given array
 */
func main() {
	arr := []int{45, 90, 80, 23, -5, 45, -6}
	fmt.Println(arr)
	x := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		temp := arr[i]
		arr[i] = x
		if x < temp {
			x = temp
		}
	}

	fmt.Println(arr)
	return
}
