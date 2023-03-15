package main

import "fmt"

/*
*This function is used to find the determinant value of 5 X 5 matrix
 */
func main() {
	var matrix = [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	n := 5
	result := Determinant(matrix, n)
	fmt.Println("The determinant of the matrix is", result)
	return
}
func Determinant(matrix [5][5]int, n int) int {

	var array [5][5]int
	dtrmnt := 0
	sign := -1

	if n == 3 {
		for i := 0; i < n; i++ {
			dtrmnt = dtrmnt + (matrix[0][i] * (matrix[1][(i+1)%3]*matrix[2][(i+2)%3] - matrix[1][(i+2)%3]*matrix[2][(i+2)%3]))
		}
		return dtrmnt
	} else {
		for i := 0; i < n; i++ {
			row := 0
			column := 0
			for j := 0; j < n; j++ {
				for k := 0; k < n; k++ {
					if j != 0 && k != i {
						array[row][column] = matrix[j][k]
						column++
					}
					if column > n-2 {
						row++
						column = 0
					}
				}
			}
			dtrmnt = dtrmnt + sign*(matrix[0][i]*Determinant(array, n-1))
			sign = -1 * sign
		}
	}
	return dtrmnt
}
