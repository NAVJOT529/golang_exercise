package main

import "fmt"

/*
*This function is used to sort the elements of the 4 X 4 matrix
 */
func main() {
	var matrix = [4][4]int{
		{3, 14, 1, 11},
		{10, 6, 9, 8},
		{7, 5, 4, 12},
		{13, 2, 16, 15},
	}
	var arr [16]int
	k := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			arr[k] = matrix[i][j]
			k++
		}
	}
	for i := 0; i < 15; i++ {
		for j := i; j < 15; j++ {
			if arr[i] > arr[j+1] {
				swap(&arr[i], &arr[j+1])
			}
		}
	}
	m := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrix[i][j] = arr[m]
			fmt.Printf("%d\t", matrix[i][j])
			m++
		}
		fmt.Println()
	}
	return
}

/* This function is used for swap the values */
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
