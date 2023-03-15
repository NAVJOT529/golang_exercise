package main

import "fmt"

/*
*This function is used to   add first seven terms of the following series using a for loop: 1/1!+2/2!+3/3! ……
 */
func main() {
	totalResult := 0.0
	factorial := 1.0
	/* using for loop to calculate the seven terms of the series*/
	for i := 1; i <= 7; i++ {
		/* changing the value of i from integer to float because we need the result in float value*/
		factorial = factorial * float64(i)
		result := float64(i) / factorial
		totalResult = totalResult + result
	}
	/* output of the sum of first seven terms of the series*/
	fmt.Println(totalResult, "is the sum of first seven terms of the series")
	return
}
