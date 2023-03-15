package main

import "fmt"

// this function is used to check the library fine of student for returning the book
func main() {
	var day int
	// taking user input for how many days student late to returning the book
	fmt.Println("Enter the number of days student late to return the book")
	fmt.Scanln(&day)
	// conditions for check the fine to student late for returning the book
	if day <= 5 {
		fmt.Println("The fine is 50 Paise")
	} else if day >= 6 && day <= 10 {
		fmt.Println("The fine is one rupee")
	} else if day > 10 && day <= 30 {
		fmt.Println("The fine is 5 rupees")
	} else {
		fmt.Println("Your membership is cancelled")
	}
	return
}
