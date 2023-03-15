package main

import "fmt"

/*
*this function is used to find the factorial value of any number entered through the keyboard.
 */
func main() {
	var number int
	factorial := 1
	/* 	taking the user input of a number which we want to check factorial  */
	fmt.Println("Enter any number")
	fmt.Scanln(&number)
	/* applying for loop to check factorial of number entered from keyboard */
	for i := 1; i <= number; i++ {
		factorial = factorial * i
	}
	fmt.Println("The factorial value of", number, "is", factorial)
	return
}
