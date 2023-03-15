package main

import "fmt"

/* This function is used to create a puzzle game*/
func main() {
	var matrix = [4][4]int{{1, 4, 15, 7}, {8, 10, 2, 11}, {14, 3, 6, 13}, {12, 9, 5, 0}}
	var testA = [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}
	var testB = [16]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var array [16]int
	fmt.Println("Solve the puzzle")
	k := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d  ", matrix[i][j])
			array[k] = matrix[i][j]
			k++
		}
		fmt.Println()
	}

	n := 15

	var loop, move string
	/* To move Right:A Left:D Down:W up:S */
	fmt.Println("Move 0 with the help of W,A,S,D key")
	for loop := "continue"; loop == "continue"; {
		fmt.Println("Enter your move")
		fmt.Scanln(&move)
		if move == "A" {
			if n == 0 || n == 4 || n == 8 || n == 12 {
				continue
			} else {
				Swap(&array[n], &array[n-1])
				Result(array)
				n = n - 1
			}
		}
		if move == "D" {
			if n == 3 || n == 7 || n == 11 || n == 15 {
				continue
			} else {
				Swap(&array[n], &array[n+1])
				Result(array)
				n = n + 1
			}
		}
		if move == "W" {
			if n == 0 || n == 1 || n == 2 || n == 3 {
				continue
			} else {
				Swap(&array[n], &array[n-4])
				Result(array)
				n = n - 4
			}
		}
		if move == "S" {
			if n == 12 || n == 13 || n == 14 || n == 15 {
				continue
			} else {
				Swap(&array[n], &array[n+4])
				Result(array)
				n = n + 4
			}
		}
		for i := 0; i < 16; i++ {
			if array[i] != testA[i] || array[i] != testB[i] {
				loop = "continue"
			} else {
				loop = "won"
			}
		}
		fmt.Println("enter continue if you want to continue the game")
		fmt.Scanln(&loop)
	}
	if loop == "won" {
		fmt.Println("You won the game")
	} else {
		fmt.Println("Try again next time you Break the game")
	}
	return
}

/* This function is used for swap the value */
func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

/* This function is used to print the result after every move */
func Result(array [16]int) [4][4]int {
	k := 0
	var matrix [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrix[i][j] = array[k]
			fmt.Printf("%d ", matrix[i][j])
			k++
		}
		fmt.Println()
	}
	return matrix
}
