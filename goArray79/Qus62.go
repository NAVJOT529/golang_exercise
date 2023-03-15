/* 62. Write a golang program to find the equilibrium indices from a given array of integers.   */
package main

import (
	"fmt"
)

/* This func is used to find the equilibrium indices from a given array of integers */
func main() {
	arr := []int{-7, 1, 5, 2, -4, 3, 0}
	fmt.Println(arr)
	totalSum := 0
	for i := range arr {
		totalSum += arr[i]
	}
	if totalSum == 0 {
		sum := 0
		for i := range arr {
			element := arr[i]
			if totalSum-sum-element == sum {
				fmt.Println("The equilibrium indices from a given array of integers", i)
			}
			sum += element
		}
	} else {
		fmt.Println("Their have no equilibrium indices in  given array of integers")
	}
	return
}
