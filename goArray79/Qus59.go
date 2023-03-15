/* 59. Write a golang program to find maximum product of two integers in a given array of integers.  */
package main

import "fmt"

/*
* This function is used to find the maximum product of two integers in a given array of integers
 */
func main() {
	var arr = []int{2, 3, 5, 7, -7, 5, 8, -5}
	fmt.Println(arr)
	max := arr[0]
	a := 0
	b := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]*arr[j] > max {
				a = arr[i]
				b = arr[j]
				max = arr[i] * arr[j]
			}
		}
	}
	fmt.Println("The product of", a, "*", b, "=", max, "is the maximum product of the array")
	return
}
