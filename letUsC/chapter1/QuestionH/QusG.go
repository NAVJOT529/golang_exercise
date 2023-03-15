package main

import "fmt"

// this function is used for calculating the sum of five digit number entered by user
func main() {
	var num int
	// number entered by the user
	fmt.Println("Enter the five digit number")
	fmt.Scanln(&num)
	// using formula for target all digits present in the number entered by user
	num5 := num % 10
	num4 := (num / 10) % 10
	num3 := (num / 100) % 10
	num2 := (num / 1000) % 10
	num1 := (num / 10000) % 10
	// output of sum of all digits present in the number entered by user
	fmt.Printf("The sum of all digit of the number you entered is %d", num1+num2+num3+num4+num5)
	return
}
