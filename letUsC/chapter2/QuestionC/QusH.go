package main

import "fmt"

// this function is used to find the absolute value of the number entered by the user
func main() {
	var num int
	// taking the user input of a number
	fmt.Println("Enter any number")
	fmt.Scanln(&num)
	// this condition is used to find the absolute value of the number entered by the user
	if num < 0 {
		fmt.Printf("The absolute value of the number is %d", (-1 * num))
	} else if num >= 0 {
		fmt.Printf("The absolute value of the number is %d", num)
	}
	return
}
