/* 17. Write a golang program to find the second largest element in an array. */
package main

import (
	"fmt"
	"sort"
)

/*
*This function is used to find the second largest element in an array
 */
func main() {
	var arr = []int{-90, 100, -20, 75, 30, 48, 69, 70, 82, 82}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)
	var i int
	for i = len(arr) - 1; arr[i] == arr[len(arr)-1]; i-- {
		fmt.Println(arr[i])
	}
	fmt.Println("Second largest element of array is", arr[i])
	return
}
