/* Write a golang program to find all triplets equal to a given sum in a unsorted array of integers.*/
package main

import (
	"fmt"
	"sort"
)

/*
*This func is used to find all triplets equal to A given sum in a unsorted array
 */
func main() {
	arr := []int{1, 6, 3, 0, 8, 4, 1, 7}
	fmt.Println("The given array is:", arr)
	sum := 7
	var array []int
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == sum {
					array = append(array, arr[i], arr[j], arr[k])
					sort.Ints(array)
					fmt.Println(array)
					array = nil
				}
			}
		}
	}
	return
}
