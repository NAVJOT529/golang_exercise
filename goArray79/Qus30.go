/* 30. Write a golang program to check if an array of integers without 0 and -1. */
package main

import "fmt"

/*
*This function is used to check if an array of integers without o and -1
 */
func main() {
	var arr = []int{2, 7, 8, 9, 4, 5, 0, -1, 4, 5, 5}
	fmt.Println(arr)
	zero := 0
	minusOne := 0
	for i := range arr {
		if arr[i] == 0 {
			zero++
		}
		if arr[i] == -1 {
			minusOne++
		}
	}
	if zero >= 1 && minusOne >= 1 {
		fmt.Println("Array contains 0 or -1 elements")
	} else {
		fmt.Println("Array doesn't contains 0 or -1 elements")
	}
	return
}
