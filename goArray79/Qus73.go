/* Write a golang program to sort a given array of distinct integers where all its numbers are sorted except two numbers. */
package main

import (
	"fmt"
	"sort"
)

/*
*This func is used to sort the given array
 */
func main() {
	arr := []int{5, 0, 1, 2, 3, 4, -2}
	fmt.Println("The given array is:", arr)
	sort.Ints(arr)
	fmt.Println("The sorted array is:", arr)
	return
}
