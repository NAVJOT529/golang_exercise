package main

import (
	"fmt"
	"math"
)

/*
*This function is used to find the norm of a matrix. The norm is defined as the square root of the sum of squares of all elements in the matrix.
 */
func main() {
	var matrix = [4][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}
	sum := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", matrix[i][j])
			sum = sum + int(math.Pow(float64(matrix[i][j]), 2))
		}
		fmt.Println()
	}
	fmt.Println("The norm of the matrix is ", int(math.Sqrt(float64(sum))))
	return
}
