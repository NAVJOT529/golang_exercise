package main

import (
	"fmt"
	"math"
)

/*
*This function is used to take the user input of any five integers
 */
func main() {
	var num1, num2, num3, num4, num5 int
	/* taking the user input of 5 integers */
	fmt.Println("Enter any five integers")
	fmt.Scanln(&num1, &num2, &num3, &num4, &num5)
	/* calling the functions*/
	sum := CalSum(num1, num2, num3, num4, num5)
	average := CalAverage(num1, num2, num3, num4, num5)
	standard := CalStandard(num1, num2, num3, num4, num5)
	/* printing the result*/
	fmt.Println("The sum of five integers is ", sum)
	fmt.Println("The average of five integers is ", average)
	fmt.Println("The standard deviation of five integers is ", standard)
	return
}

/* This function is used to calculate the sum of the integers */
func CalSum(num1, num2, num3, num4, num5 int) int {
	sum := num1 + num2 + num3 + num4 + num5
	return sum
}

/* This function is used to calculate the average of the integers */
func CalAverage(num1, num2, num3, num4, num5 int) int {
	average := (num1 + num2 + num3 + num4 + num5) / 5
	return average
}

/* This function is used to calculate the standard deviation of the integers */
func CalStandard(num1, num2, num3, num4, num5 int) int {
	average := (num1 + num2 + num3 + num4 + num5) / 5
	sum := math.Pow(float64(num1-average), 2) + math.Pow(float64(num2-average), 2) + math.Pow(float64(num3-average), 2) + math.Pow(float64(num4-average), 2) + math.Pow(float64(num5-average), 2)
	return int(math.Sqrt(sum / 5))
}
