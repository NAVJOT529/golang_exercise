package main

import "fmt"

/*
*This function is used to produce a this output
A B C D E F G F E D C B A
A B C D E F   F E D C B A
A B C D E       E D C B A
A B C D           D C B A
A B C               C B A
A B                   B A
A                       A
*/
func main() {
	/* This for loop is dedicated to every row present in the Output*/
	for row := 0; row <= 6; row++ {
		/* This for loop create a section from A  in every row*/
		for part1 := 65; part1 <= 71-row; part1++ {
			fmt.Printf("%c ", part1)
		}
		/* This for loop creates space between first section and second section from row 2*/
		for space := 1; space <= row*2-1; space++ {
			fmt.Printf("  ")
		}
		/* This for loop creates second reverse section of A in every row*/
		for part2 := 71 - row; part2 >= 65; part2-- {
			if part2 == 71 {
				continue
			} else {
				fmt.Printf("%c ", part2)
			}
		}
		/* here every row take a new line */
		fmt.Printf("\n")
	}
	return
}
