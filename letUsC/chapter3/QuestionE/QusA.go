package main

import "fmt"

/*
* This function is used to print all prime numbers from 1 to 300.
 */
func main() {
	/* applying loop to check all the prime number from 1 to 300*/
	for num := 1; num < 300; num++ {
		mark := 0
		/* this loop is apply to check the prime number */
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				mark = 1
				break
			}
		}
		/* output for all the prime numbers*/
		if mark == 0 && num != 1 {
			fmt.Println("The num", num, "is prime number")
		}
	}
	return
}
