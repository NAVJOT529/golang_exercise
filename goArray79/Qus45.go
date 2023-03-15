/* 45. Write a golang program to cyclically rotate a given array clockwise by one.  */
package main

import "fmt"

/*
This function is used to rotate a given array clockwise by one
*/
func main() {
	var array = []int{10, 20, 30, 40, 50, 60}
	fmt.Println(array)
	var arr []int
	for i := len(array) - 1; i >= 0; i-- {
		val := array[i]
		arr = append(arr, val)
	}

	for i := 0; i < len(arr)-1; i++ {
		swap(&arr[i], &arr[i+1])
	}

	var reversed []int
	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		reversed = append(reversed, val)
	}
	fmt.Println(reversed)
	return
}

/* This function is used for swap the values */
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
