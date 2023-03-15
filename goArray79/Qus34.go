/* 34. Write a golang  program to find the length of the longest consecutive elements sequence from a given unsorted array of integers. Sample array: [49, 1, 3, 200, 2, 4, 70, 5] The longest consecutive elements sequence is [1, 2, 3, 4, 5], therefore the program will return its length 5.*/
package main

import (
	"fmt"
	"sort"
)

/*
*This function is used to find the length of longest consecutive elements sequence in the array
 */
func main() {
	var arr = []int{49, 1, 3, 200, 2, 4, 70, 5}
	fmt.Println("The given array is:", arr)
	sort.Ints(arr)
	fmt.Println(arr)
	var array []int
	array = append(array, arr[0])
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+1 == arr[i+1] {
			array = append(array, arr[i+1])
		}
	}
	fmt.Println("longest consecutive elements sequence is :", array, "With the length of", len(array))
	return
}
