package main

import "fmt"

// this function is use to find new number by adding 1 in each digit of the number enter by the user
func main() {
	var num int
	// taking a user input of five digit number
	fmt.Println("Enter a five digit number")
	fmt.Scanln(&num)
	// using formula to target every digit of number and adding one in each digit
	num5 := ((num % 10) + 1) % 10
	num4 := (((num / 10) % 10) + 1) % 10
	num3 := (((num / 100) % 10) + 1) % 10
	num2 := (((num / 1000) % 10) + 1) % 10
	num1 := (((num / 10000) % 10) + 1) % 10
	// now finding the new number
	newNumber := num1*10000 + num2*1000 + num3*100 + num4*10 + num5
	// output of the new number after adding one in each digit of the number
	fmt.Printf("The new number by adding one in each digit of the number is %d", newNumber)
	return
}
