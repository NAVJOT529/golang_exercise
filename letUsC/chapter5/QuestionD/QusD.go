package main

import "fmt"

/*
*This function is used to take a user input of the year
 */
func main() {
	var year int
	/* Taking the user input for a year */
	fmt.Println("Enter any year")
	fmt.Scanln(&year)
	LeapYearOrNot(year)
	return
}

/* This function is used to find the year entered by the user is leap yea or not*/
func LeapYearOrNot(year int) {
	if year%4 == 0 || year%400 == 0 && year%100 != 0 {
		fmt.Println("The year", year, "is a leap year")
	} else {
		fmt.Println("The year", year, "is not a leap year")
	}
	return
}
