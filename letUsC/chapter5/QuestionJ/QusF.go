package main

import (
	"fmt"
	"math"
)

/*
*This function is used to take the user input of value of x and number of terms in the series
 */
func main() {
	fmt.Println("Enter the value of x")
	var x float64
	fmt.Scanln(&x)
	fmt.Println("Enter the number of terms")
	var n int
	fmt.Scanln(&n)
	result := Evaluate(x, n)
	fmt.Println("The value of sin(", x, ")is :", result)
}

/* This function is used to evaluate the given series  */
func Evaluate(x float64, n int) float64 {
	sum := 0.0
	power := 1.0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum = sum - float64(math.Pow(x, float64(power))/Factorial(power))
			power = power + 2
		} else {
			sum = sum + float64(math.Pow(x, float64(power)))/Factorial(power)
			power = power + 2
		}
	}
	return float64(sum)
}

/* This function is used to find the factorial */
func Factorial(x float64) float64 {
	var facto float64
	if x == 1 {
		return 1
	} else {
		facto = x * Factorial(x-1)
	}
	return facto
}
