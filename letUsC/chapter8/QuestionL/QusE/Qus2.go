package main

import "fmt"

/*
*This function is used to multiply any 3 X 3 matrix
 */
func main() {
	var one = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var two = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var mult [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			mult[i][j] = 0
			for k := 0; k < 3; k++ {
				mult[i][j] = mult[i][j] + one[i][k]*two[j][i]
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", mult[i][j])
		}
		fmt.Println("")
	}
	return
}
