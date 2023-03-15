/* Write a Java program to find maximum difference between two elements in a given array of integers such that smaller element appears before larger element */
package main

import "fmt"

/*
*This function is used to find maximum difference between two elements in a given array of integers
 */
func main() {
	arr := []int{2, 3, 1, 7, 9, 5, 11, 3, 5}
	fmt.Println("Given array of integers :", arr)
	maxDiff := 1
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]-arr[i] > maxDiff {
				maxDiff = arr[j] - arr[i]
			}
		}
	}
	fmt.Println("The maximum difference between two elements in a given array :", maxDiff)
	return
}
