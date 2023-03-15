package main

import (
	"fmt"
	"math"
)

/*
*This function is used to read 10 sets of p, r, n & q and calculate the amount corresponding as.
 */
func main() {
	var time, rate, year, principle float64
	for set := 1; set <= 10; set++ {
		fmt.Println("Enter the time rate, annual rate, year, and principle of the compound")
		fmt.Scanln(&time, &rate, &year, &principle)
		amount := principle * math.Pow((1+rate/time), rate*time)
		fmt.Println("amount =", amount)
	}
	return
}
