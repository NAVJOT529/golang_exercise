/* Write a golang program to rearrange a given array of unique elements such that every second element of the array is greater than its left and right elements. */
package main

import "fmt"

/*
*This function is used to arrange the unique elements
 */
func main() {
	arr := []int{1, 2, 4, 9, 5, 3, 8, 7, 10, 12, 14}
	fmt.Println(arr)
	for i := 1; i < len(arr); i += 2 {
		if arr[i-1] > arr[i] {
			swap(&arr[i], &arr[i-1])
		}
		if i+1 < len(arr) && arr[i+1] > arr[i] {
			swap(&arr[i], &arr[i+1])
		}
	}
	fmt.Println(arr)
	return
}

/* This function is used to swap the elements */
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
