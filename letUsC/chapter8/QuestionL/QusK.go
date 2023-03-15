package main

import "fmt"

/*
*This function is used to find the square matrix is symmetric
 */
func main() {
	var matrix = [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	var trans [4][4]int
	fmt.Println("transpose of matrix")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d\t", matrix[i][j])
			trans[j][i] = matrix[i][j]
		}
		fmt.Println()
	}
	mark := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if trans[i][j] == matrix[i][j] {
				mark = 0 + mark
			} else {
				mark = 1 + mark
			}
		}
	}
	if mark == 0 {
		fmt.Println("The square matrix is symmetric")
	} else {
		fmt.Println("The square matrix is not symmetric")
	}
	return
}
