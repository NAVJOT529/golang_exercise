/* 15. Write a golang program to find the common elements between two arrays of integers. */
package main

import "fmt"

/*
*This function is used to find the common factor in an integer array
 */
func main() {
	var array = []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	var arr = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	var common []int
	fmt.Println("multiplication table of 2   ", array)
	fmt.Println("multiplication table of 5   ", arr)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(arr); j++ {
			if array[i] == arr[j] {
				common = append(common, array[i])
			}
		}
	}
	fmt.Println("common factor 2 and 5 multiplication table", common)
	return
}
