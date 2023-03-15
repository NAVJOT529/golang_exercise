/* 19. Write a go lang program to add two matrices of the same size. */
package main

import "fmt"

/*
*This function is used to find the sum of two matrices of the same size
 */
func main() {
	array := [4][4]int{{10, 20, 30, 40}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	arr := [4][4]int{{10, 20, 30, 40}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	var sum [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum[i][j] = array[i][j] + arr[i][j]
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d  ", sum[i][j])
		}
		fmt.Println()
	}
	return
}
