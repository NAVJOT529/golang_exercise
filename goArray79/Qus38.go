/* Write a golang program to get the majority element from a given array of integers containing duplicates. Majority element: A majority element is an element that appears more than n/2 times where n is the size of the array.*/
package main

import "fmt"

/*
*This function is used to find the majority element from the given array
 */
func main() {
	var arr = []int{1, 6, 6, 5, 7, 4, 1, 7, 7, 7, 7, 7, 7, 7, 2}
	fmt.Println(arr)
	count := 1
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count > len(arr)/2 {
			fmt.Println("The majority element in the array is ", arr[i])
			break
		}
	}
	return
}
