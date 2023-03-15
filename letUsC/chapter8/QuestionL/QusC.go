package main

import "fmt"

/*
*This function is used to obtained the transpose of 4 X 4 matrix
 */
func main() {
	var matrix [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println("Enter the element of the matrix")
			fmt.Scanln(&matrix[i][j])
		}
	}
	fmt.Println("The Transpose of the matrix")
	var transpose [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", matrix[i][j])
			transpose[j][i] = matrix[i][j]
		}
		fmt.Println()
	}
	fmt.Println("is")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", transpose[i][j])
		}
		fmt.Println()
	}
	return
}
