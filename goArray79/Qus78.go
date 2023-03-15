/*  Write a golang program that checks an array is negative dominant or not. If the array is negative dominant return true otherwise false. */
package main

import "fmt"

/*
*This function is used to Print the output if the array is negative dominant or not
 */
func main() {
	arr := []int{1, 2, -5, -4, 3, -6}
	fmt.Println(arr)
	fmt.Println("This function is negative dominant: ", Check(arr))
	return
}

/* This func is used to check the arr is negative dominant or not  */
func Check(arr []int) bool {
	negative := 0
	positive := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative++
		}
		if arr[i] > 0 {
			positive++
		}
	}
	if negative > positive {
		return true
	}
	return false
}
