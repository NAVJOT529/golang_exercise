/* 16. Write a golang program to remove duplicate elements from an array. */
package main

import "fmt"

/*
*This function is used to remove duplicate elements from an array
 */
func main() {
	array := []int{1, 2, 2, 5, 2, 8, 8, 6}
	fmt.Println("Original array:\t", array)
	var dup []int
	encountered := map[int]bool{}
	for i := range array {
		encountered[array[i]] = true
	}
	for key := range encountered {
		dup = append(dup, key)
	}
	fmt.Println("After remove duplicate values :\t", dup)
	return
}
