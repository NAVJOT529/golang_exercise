/* . Write a golang program that checks whether an array of integers alternates between positive and negative values.  */
package main

import "fmt"

/*
*This function is used to Print if array is a alternates between positive and negative or not
 */
func main() {
	arr := []int{-1, 2, -5, 6, -4, -3, -6}
	fmt.Println(arr)
	fmt.Println(Check(arr))
	return
}

/* This function is used to check if the elements of array are alternates positive and negative values */
func Check(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > 0 && arr[i+1] > 0 {
			return false
		} else if arr[i] < 0 && arr[i+1] < 0 {
			return false
		}
	}
	return true
}
