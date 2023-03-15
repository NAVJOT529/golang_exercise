/* 48. Write a golang program to arrange the elements of a given array of integers where all negative integers appear before all the positive integers. */
package main

import "fmt"

/*
* This func is used to arrange the elements of the array where all negative integers appear before all the positive integers.
 */
func main() {
	var arr = []int{-4, 8, 6, -5, 6, -2, 1, 2, 3, -11}
	fmt.Println("The give array is ", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > 0 {
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
