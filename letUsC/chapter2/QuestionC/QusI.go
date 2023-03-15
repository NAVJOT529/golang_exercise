package main

import "fmt"

// this function is used to check the area of rectangle is greater than perimeter or not
func main() {
	var length, breadth int
	// taking the user input for length and breadth of rectangle
	fmt.Println("Enter the value of length and breadth of rectangle")
	fmt.Scanln(&length, &breadth)
	// condition is used to check the area of rectangle is greater than perimeter or not
	// calculating area and perimeter in if condition
	if length*breadth > 2*(length+breadth) {
		fmt.Println("The Area of rectangle is greater than its Perimeter")
	} else {
		fmt.Println("Perimeter of rectangle is greater than its Area")
	}
}
