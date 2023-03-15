/* 26. Write a golang program to move all 0's to the end of an array. Maintain the relative order of the other (non-zero) array elements. */
package main

import "fmt"

/*
*This function is used to move all 0's to the end of a array\
 */
func main() {
	var arr = []int{10, 1, 0, 0, 2, 0, 3, 4, 0, 5}
	fmt.Println("The give array is ", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr)-1; j++ {
			if arr[i] == 0 {
				swap(&arr[i], &arr[j+1])
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
