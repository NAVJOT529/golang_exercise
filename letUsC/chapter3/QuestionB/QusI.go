package main

import "fmt"

/*
* This function  is used to find the range of a set of numbers.
 */
func main() {
	var num int
	max := 0
	min := 0
	another := "yes"
	for another == "yes" {
		fmt.Println("Enter any number")
		fmt.Scanln(&num)
		if num <= min {
			min = num
		}
		if num >= max {
			max = num
		}
		fmt.Println("Do you want to enter another number? (yes/no)")
		fmt.Scanln(&another)
	}
	fmt.Println("The range of set of numbers is:", (max - min))
	return
}
