/* 51. Write a golang program to separate 0s on left side and 1s on right side of an array of 0s and 1s in random order.  */
package main

import "fmt"

/*
*This func is used to separate 0s on the left side and 1s on the right side of the array
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
