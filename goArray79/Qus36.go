/* Write a golang program to find all the unique triplets such that sum of all the three elements [x, y, z (x ≤ y ≤ z)] equal to a specified number. Sample array: [1, -2, 0, 5, -1, -4] Target value: 2. */
package main

import "fmt"

/*
*This function is used to find the unique triplets in the array
 */
func main() {
	var arr = []int{1, -2, 0, 5, -1, -4}
	fmt.Println(arr)
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 2 {
					fmt.Println("unique triples in the array is ", arr[i], arr[j], arr[k], "With the sum of 2")
				}
			}
		}
	}
	return
}
