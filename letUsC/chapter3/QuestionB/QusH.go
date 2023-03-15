package main

import "fmt"

/*
*This function is used to find the octal equivalent of the entered number.
 */
func main() {
	var num int
	term := 0
	octal := 0
	/* taking the user input to get the number to find its octal equivalent*/
	fmt.Println("Enter any number: ")
	fmt.Scanln(&num)
	/* applying a loop to find octal */
	for num != 0 {
		module := num % 8
		term = term * 10
		if term == 0 {
			term = 1
		}
		octal = term*module + octal
		num = num / 8
	}
	fmt.Println("The octal equivalent of the entered number is: ", octal)
	return
}
