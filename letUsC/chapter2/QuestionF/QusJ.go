package main

import "fmt"

// this function is to process the customer order
func main() {
	var order int
	var credit string
	stock := 1000
	// taking the user input of order and his credits
	fmt.Println("Enter your order Quantity and Credits")
	fmt.Scanln(&order, &credit)
	// applying condition to check the customer order is available or not
	if order <= stock && credit == "ok" {
		fmt.Println("supply has requirements")
	} else if order <= stock && credit == "notok" {
		fmt.Println("Your credits are not ok")
	} else if order > stock && credit == "ok" {
		fmt.Println("The balance will be shipped")
	}
	return
}
