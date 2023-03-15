package main

import "fmt"

/*
*This function is used to pickup the largest value of a element from a 5 x 5 matrix
 */
func main() {
	var matrix [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("Enter the element of the matrix")
			fmt.Scanln(&matrix[i][j])
		}
	}
	max := matrix[0][0]
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] >= max {
				max = matrix[i][j]
			}
		}
	}
	fmt.Println("The largest element is : ", max)
	return
}
