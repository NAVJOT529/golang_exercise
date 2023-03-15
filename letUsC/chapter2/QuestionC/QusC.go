package main

import "fmt"

// this function is used to check the the year entered by user is leap year or not
func main() {
	var year int
	// taking user input of year which user want to check leap year or not
	fmt.Println("Enter any year")
	fmt.Scanln(&year)
	// this condition is used to check year enter by user is leap year or not
	if (year%4 == 0 || year%400 == 0) && year%100 != 0 {
		fmt.Println("This is a leap year")
	} else {
		fmt.Println("This is not a leap year")
	}
	return
}
