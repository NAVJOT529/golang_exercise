package main

import "fmt"

/*
*This function is used to take user input of any number and print its factorial value
 */
func main() {
	var num int
	/* taking user input */
	fmt.Println("Enter any number")
	fmt.Scanln(&num)
	/* calling function */
	fact := Factorial(num)
	/* output of factorial of the number entered by the user */
	fmt.Println("The factorial value of the number you entered is: ", fact)
	return
}

/* This function is used to calculate the factorial value of integer entered through the keyboard in main function */
func Factorial(num int) int {
	fact := 1
	for i := 1; i <= num; i++ {
		fact = fact * i
	}
	return fact
}
