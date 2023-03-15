/* 24. Write a golang program to find a missing number in an array. */
package main

import "fmt"

/*
*This function is used to find the missing number in an array
 */
func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	fmt.Println("The given array :", array)
	expectedNumber := ((len(array) + 1) * (len(array) + 2)) / 2
	sum := 0
	for i := range array {
		sum = sum + array[i]
	}
	fmt.Println("The missing number in the array is ", expectedNumber-sum)
	return
}
