package main

import (
	"fmt"
	"math"
)

/*
*This function is used to  print a Pascal's triangle
 */
func main() {
	var module int
	/* This loop is used to target every row in Pascal's triangle*/
	for row := 1; row <= 5; row++ {
		/* This loop is used to make space  */
		for space := 1; space <= 5-row; space++ {
			fmt.Print("  ")
		}
		/* This nested loop is used to make the numbers for pascal's triangle*/
		for power := row - 1; power == row-1; power++ {
			num := math.Pow(11, float64(power))
			for num != 0 {
				module = int(num) % 10
				if module == 0 {
					break
				} else {
					/* here is the output of pascal's triangle*/
					fmt.Printf("%d ", module)
					fmt.Printf("  ")
					num = num / 10
				}
			}
		}
		/* this is use to go for the new row*/
		fmt.Printf("\n")
	}
	return
}
