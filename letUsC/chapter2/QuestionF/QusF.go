package main

import "fmt"

// this function is used to check the triangle is valid or not
func main() {
	var side1, side2, side3 int
	// taking the user input for three sides of triangle
	fmt.Println("Enter the three sides of triangle side1,side2 and side3")
	fmt.Scanln(&side1, &side2, &side3)
	// applying conditions to check the triangle is valid or not a triangle is valid when greater side of triangle is equal to sum of another sides
	if side1 > side2 && side1 > side3 {
		if side1 > side2+side3 {
			fmt.Println("This is a valid triangle")
		} else {
			fmt.Println("This is a invalid triangle")
		}
	} else if side2 > side1 && side2 > side3 {
		if side2 > side1+side3 {
			fmt.Println("This is a valid triangle")
		} else {
			fmt.Println("This is a invalid triangle")
		}
	} else if side3 > side2 && side3 > side1 {
		if side3 > side1+side2 {
			fmt.Println("This is a valid triangle")
		} else {
			fmt.Println("This is a invalid triangle")
		}
	} else {
		fmt.Println("This is a invalid triangle")
	}
	return
}
