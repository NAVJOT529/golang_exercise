/* 47. Write a golang program to find the rotation count in a given rotated sorted array of integers. */
package main

import "fmt"

/*
*This function is used to count the rotation of a shorted array
 */
func main() {
	var arr = []int{35, 32, 30, 14, 18, 21, 27}
	fmt.Println(arr)
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(&arr[i], &arr[j])
				count++
			}
		}
	}
	fmt.Println("There have", count, "rotations to sort the array", arr)
}

/* This function is used for swap the values */
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
