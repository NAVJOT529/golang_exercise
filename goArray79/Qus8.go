/* 8. Write a golang program to copy an array by iterating the array. */
package main

import "fmt"

/*
*This function is used to copy an array into another
 */
func main() {
	var arr = []int{24, 36, 28, 29, 30, 32, 47, 52, 49, 68}
	fmt.Println("Original array: ", arr)
	var array []int
	var val int
	for i := range arr {
		val = arr[i]
		array = append(array, val)
	}
	fmt.Println("copied array: ", array)
	return
}
