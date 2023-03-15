package main

import (
	"fmt"
	"math"
)

/*
*This function is used to print out all Armstrong numbers between 1 and 500
 */
func main() {
	var num int
	/* using for loop to check the armstrong numbers between 1 and 500*/
	for num = 1; num <= 500; num++ {
		/* to target every digit in the number */
		a := (num / 100) % 10
		b := (num / 10) % 10
		c := num % 10
		/* applying condition for check the number is armstrong or not if this condition is true the number is armstrong else no */
		if float64(num) == math.Pow(float64(a), 3)+math.Pow(float64(b), 3)+math.Pow(float64(c), 3) {
			fmt.Println("The number", num, "is a armstrong number")
		}
	}
}
