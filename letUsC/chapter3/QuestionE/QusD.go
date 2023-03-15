package main

import "fmt"

/*
*This function is used to generate all combinations of 1, 2 and 3.
 */
func main() {
	/* using for loop to generate all combinations of 1 2 and 3*/
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == j {
				/* if i == j then that condition will continue because we don't need same numbers*/
				continue
			} else {
				for k := 1; k <= 3; k++ {
					if i != j && j != k && k != i {
						/* when number 1 2 3 make any combinations that result will print from here*/
						fmt.Println(i, j, k)
					}
				}
			}
		}
	}
	return
}
