/*Write a golang program to determine whether numbers in an array can be rearranged so that each number appears exactly once in a consecutive list of numbers. Return true otherwise false */
package main

import (
	"fmt"
	"sort"
)

/*
*This function is used to print the output for if array is a consecutive list of numbers or not
 */
func main() {
	arr := []int{1, 2, 5, 0, 4, 3, 6}
	fmt.Println(arr)
	fmt.Println(Check(arr))
	return
}

/* This function is used to check the array is a consecutive or not  */
func Check(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		if i+arr[i] == arr[i+1] {
			return true
		}
	}
	return false
}
