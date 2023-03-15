package main

import "fmt"

/*
*This function is used to print the output of the sum of first 25 natural numbers*/

func main() {
	var sum int
	num := 25
	/* Calling the function*/
	sum = CalSum(num)
	/* Output of sum of first 25 natural numbers*/
	fmt.Println("The sum of the first 25 natural number is", sum)
	return
}

/* This function is used to find the running sum of first 25 natural numbers by recursive function */
func CalSum(num int) int {
	var sum int
	if num != 0 {
		sum = num + CalSum(num-1)
	}
	return sum
}
