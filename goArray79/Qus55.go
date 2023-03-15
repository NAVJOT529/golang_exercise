/*  Write a golang program to print all sub-arrays with 0 sum present in a given array of integers. */
package main

import "fmt"

/*
*This func is used to Print all sub-arrays with 0 sum present in the array
 */
func main() {
	arr := []int{1, 3, -7, 3, 2, 3, 1, -3, -2, -2}
	fmt.Println(arr)
	array := []int{}
	for i := range arr {
		sum := 0
		array = nil
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			array = append(array, arr[j])
			if sum == 0 {
				fmt.Println("sum of the Elements of array:", array, "is 0")
			}
		}
	}
	return
}
