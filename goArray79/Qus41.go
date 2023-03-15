/* 41. Write a golang program to find smallest and second smallest elements of a given array. */
package main

import (
	"fmt"
	"sort"
)

/*
*This func is used to find the smallest and second smallest element of a array
 */
func main() {
	var arr = []int{44, 3, 3, 3, 50, 6, 88, 97, 23, 32, 3}
	fmt.Println("Given array is", arr)
	sort.Ints(arr)
	fmt.Println("After sort the array", arr)
	var i int
	for i = 0; arr[i] == arr[i+1]; i++ {
	}
	fmt.Println("First smallest element in the array is      ", arr[0], "\nThe second smallest element in the array is ", arr[i+1])
	return
}
