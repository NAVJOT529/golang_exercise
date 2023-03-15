/* 42. Write a golang program to segregate all 0s on left side and all 1s on right side of a given array of 0s and 1s. */
package main

import "fmt"

/*
*This func is used to segregate all 0s of the array to the left side and all ones to the right
 */
func main() {
	var arr = []int{0, 1, 0, 0, 1, 0, 1, 1, 0, 0}
	fmt.Println("The give array is ", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == 1 {
				swap(&arr[i], &arr[j])
			}
		}
	}
	fmt.Println("The new array is :", arr)
	return
}

/* This function is used for swap the values */
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
