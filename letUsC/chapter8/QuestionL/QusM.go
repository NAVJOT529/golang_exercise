package main

import "fmt"

/*
*This function is used to shift the elements of array p[5] circularly
 */
func main() {
	var matrix [4][5]int
	var p = [5]int{15, 30, 28, 19, 61}
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = p[j]
		}
		swap(&p[0], &p[3])
		swap(&p[1], &p[4])
		swap(&p[2], &p[0])
		swap(&p[1], &p[2])
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	return
}

/* This function is used to swap the elements */
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
