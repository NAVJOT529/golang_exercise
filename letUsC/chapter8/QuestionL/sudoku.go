package main

import "fmt"

/*
*This function is used to check the element entered by user in the array is in sudoku form or not
 */
func main() {
	var matrix [9][9]int
	fmt.Println("Enter the elements of sudoku one by one")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
		fmt.Println("Next Row")
	}
	mark := 1
	sum := 0
	add := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sum = sum + matrix[i][j]
			add = add + matrix[j][i]
		}
		if sum != 45 || add != 45 {
			fmt.Println("matrix is not in sudoku form")
			mark = 0
			break
		} else {
			sum = 0
			add = 0
		}
	}
	if mark != 0 {
		for i := 0; i < 9; i = i + 3 {
			for j := 0; j < 9; j = j + 3 {
				for k := i; k <= i+2; k++ {
					for l := j; l <= j+2; l++ {
						sum = sum + matrix[k][l]
					}
				}
				if sum != 45 {
					fmt.Println("matrix is not in sudoku form")
					mark = 0
					break
				} else {
					sum = 0
				}
			}
		}
	}
	if mark != 0 {
		fmt.Println("matrix is in sudoku form you won the game")
	}
	return
}
