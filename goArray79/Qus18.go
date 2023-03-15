/* 18. Write a golang program to find the second smallest element in an array.  */
package main

import (
	"fmt"
	"sort"
)

/*
*This function is used to find the smallest element in the array
 */
func main() {
	var arr = []int{50, 20, 20, 100, 20, 75, 30, 48, 69, 70, 82}
	fmt.Println("Original array is:", arr)
	sort.Ints(arr)
	fmt.Println(arr)
	i := 0
	for ; arr[i] == arr[i+1]; i++ {
	}
	fmt.Println("Second smallest element of the array is:", arr[i+1])
	return
}
