/* 44. Write a golang program to count the number of possible triangles from a given unsorted array of positive integers.Note: The triangle inequality states that the sum of the lengths of any two sides of a triangle must be greater than or equal to the length of the third side. */
package main

import "fmt"

/*
*This function is used to count the possible triangle into the array
 */
func main() {
	var arr = []int{6, 7, 9, 16, 25, 12, 30, 40}
	fmt.Println(arr)
	count := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j] >= arr[k] {
					count++
				}
			}
		}
	}
	fmt.Println("Their have", count, "Possible triangles into the array")
	return
}
