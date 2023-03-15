package main

import "fmt"

// this function is used to calculate the sum of first and last digit from a number entered by user
func main() {
	var num int
	// number entered by the user
	fmt.Println("Enter the four digit number")
	fmt.Scanln(&num)
	// using formula for target first and last digit present in the number entered by user
	num4 := num % 10
	num1 := (num / 1000) % 10
	// output of sum of first and last digit of number entered by user
	fmt.Printf("The sum of first and last digit of the number you entered is %d", num1+num4)
	return
}
