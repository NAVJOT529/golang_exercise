package main

import "fmt"

/*
* This function is used to take user input for a year
 */
func main() {
	var year int
	fmt.Println("Enter any year")
	fmt.Scanln(&year)
	roman(year)
	return
}

/* This function is used to convert the given year into its roman equivalent*/
func roman(year int) {
	if year >= 1000 {
		fmt.Printf("m")
		roman(year - 1000)
	} else if year >= 500 {
		fmt.Printf("d")
		roman(year - 500)
	} else if year >= 100 {
		fmt.Printf("c")
		roman(year - 100)
	} else if year >= 50 {
		fmt.Printf("l")
		roman(year - 50)
	} else if year >= 10 {
		fmt.Printf("x")
		roman(year - 10)
	} else if year >= 5 {
		fmt.Printf("v")
		roman(year - 5)
	} else if year >= 1 {
		fmt.Printf("i")
		roman(year - 1)
	}
	return
}
