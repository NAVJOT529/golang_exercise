/* 54. Write a Java program to check if a given array contains a sub array with 0 sum. */
package main

import (
	"fmt"
)

/*
*This function is used to check if a given array contains a sub array with 0 sum
 */
func main() {
	var arr = []int{1, 2, -3, 4, 5, 6}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			fmt.Println("Sub array exist with the sum of 0")
			return
		}
	}
	array := []int{}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum == 0 || Prefix(array, sum) == sum {
			fmt.Println("Sub array exist with the sum of 0")
			return
		}
		array = append(array, sum)
	}
	fmt.Println("Sub array doesn't exist with the sum of 0")
	return
}

/* This function is used to check The Prefix sum*/
func Prefix(array []int, sum int) int {
	for i := range array {
		if array[i] == sum {
			return sum
		}
	}
	return 0
}

/* Either a prefix sum repeats or
Or prefix sum becomes 0.
Since prefix sum 1 repeats, we have a subArray with 0 sum. */
