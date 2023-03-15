package main

import "fmt"

/*
*This function is used to produce a triangular shape with numbers 1 to 10
 */
func main() {
	print := 1
	/* This loop is dedicated to the row in the triangular shape*/
	for row := 1; row <= 4; row++ {
		/* this loop is dedicated to the space in the triangular shape*/
		for space := 1; space <= 4-row; space++ {
			fmt.Print("  ")
		}
		/* this nested loop will print the number from 1 to 10 one by one in every row  */
		for part1 := 1; part1 <= row; part1++ {
			for part2 := part1; part2 <= part1; part2++ {
				fmt.Printf("%d ", print)
				print++
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
	return
}
