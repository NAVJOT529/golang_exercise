/* 46. Write a golang program to check whether there is a pair with a specified sum of a given sorted and rotated array. */
package main

import (
	"fmt"
)

/*
*This function is used to find the specified sum of a given sorted array
 */
func main() {
	var arr = []int{6, 8, 8, 9, 10, 11, 15}
	fmt.Println(arr)
	sum := 16
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Println("There have a pair of ", arr[i], "and", arr[j], " with the specified sum ", sum)
				break
			}
		}
	}
	return
}
