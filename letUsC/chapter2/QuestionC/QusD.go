package main

import "fmt"

// this function is used to check the day on first january of the year enter by the user
func main() {
	var year int
	var givenYear int = 1900
	// taking user input of year which user want to check
	fmt.Println("Enter any year")
	fmt.Scanln(&year)
	// checking the day of the year
	totalYear := year - givenYear
	leapYear := (totalYear / 4) + (totalYear / 400) - (totalYear / 100)
	remainingYear := totalYear - leapYear
	totalDays := (remainingYear*365 + leapYear*366) + 1
	day := totalDays % 7
	// these conditions is used to check the day of 1st january of the year input by user
	if day == 0 {
		fmt.Println("MONDAY")
	} else if day == 1 {
		fmt.Println("TUESDAY")
	} else if day == 2 {
		fmt.Println("WEDNESDAY")
	} else if day == 3 {
		fmt.Println("THURSDAY")
	} else if day == 4 {
		fmt.Println("FRIDAY")
	} else if day == 5 {
		fmt.Println("SATURDAY")
	} else if day == 6 {
		fmt.Println("SUNDAY")
	} else {
		fmt.Println("INVALID YEAR")
	}
	return
}
