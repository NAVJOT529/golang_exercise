package main

import "fmt"

/*
*This function is used to add two 6 X 6 matrix
 */
func main() {
	var one = [6][6]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36},
	}
	var two = [6][6]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36},
	}

	fmt.Println("The sum of the two 6X6 matrix is")
	var sum [6][6]int
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			sum[i][j] = one[i][j] + two[i][j]
			fmt.Printf("%d\t", sum[i][j])
		}
		fmt.Println()
	}
	/* var one = map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8, 8: 9, 9: 10}
	var two = map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8, 8: 9, 9: 10}
	var sum [10]int
	for i := 0; i < 10; i++ {
		sum[i] = one[i] + two[i]
	}
	fmt.Println(sum) */
	return
}
