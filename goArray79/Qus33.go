/* Write a go lang  program to remove the duplicate elements of a given array and return the new length of the array. Sample array: [20, 20, 30, 40, 50, 50, 50] After removing the duplicate elements the program should return 4 as the new length of the array. */
package main

import "fmt"

/*
*This function is used to remove duplicate elements from an array
 */
func main() {
	array := []int{20, 20, 30, 40, 50, 50, 50}
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
	fmt.Println("The new length of the array is :", len(dup))
	return
}
