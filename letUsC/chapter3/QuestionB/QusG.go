package main

import "fmt"

/*
* this function is used to Write a program to enter the numbers till the user wants and at the end it should
display the count of positive, negative and zeros entered
*/
func main() {
	another := "yes"
	var num int
	negatives := 0
	positives := 0
	zeros := 0
	/* taking a loop to user input of positive, negative and zeros and count them*/
	for another == "yes" {
		fmt.Println("Enter any number")
		fmt.Scanln(&num)
		if num < 0 {
			negatives++
		} else if num > 0 {
			positives++
		} else {
			zeros++
		}
		fmt.Println("Do you want to enter another number? (yes or no)")
		fmt.Scanln(&another)
	}
	fmt.Println("You entered", negatives, "negative numbers", positives, "positive numbers and", zeros, "zero numbers")
	return
}
