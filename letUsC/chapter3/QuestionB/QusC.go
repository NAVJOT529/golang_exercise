package main

import "fmt"

/*
*This function is used to find the value of first number raised to power second when two numbers are entered threw the keyboard
 */
func main() {
	var num1, num2 int
	var power int = 1
	/*  taking the user input to find the value of first number raised to power second */
	fmt.Println("Enter the value of first number and second number")
	fmt.Scanln(&num1, &num2)
	/* applying loop to find the value of first number raised to power second */
	for i := 1; i <= num2; i++ {
		power = power * num1
	}
	fmt.Println("The value of", num1, "raised to power", num2, "is", power)
	return
}
