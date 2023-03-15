/* 75. Write a Java program to calculate the largest gap between sorted elements of an array of integers.  */
package main

import (
	"fmt"
	"sort"
)

/*
*This func is used to calculate the largest gap between sorted elements
 */
func main() {
	arr := []int{23, -2, 45, 38, 12, 4, 6}
	fmt.Println("Original array :", arr)
	sort.Ints(arr)
	max := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] > max {
			max = arr[i+1] - arr[i]
		}
	}
	fmt.Println("Largest gap between sorted elements of the said array:", max)
	return
}
