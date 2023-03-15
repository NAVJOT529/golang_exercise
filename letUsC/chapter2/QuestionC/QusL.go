package main

import "fmt"

// this function is used to check the point lies on the x axis , y axis or at the origin
func main() {
	var x, y int
	// taking the user input for coordinates of the point
	fmt.Println("Enter the coordinates X,Y of the point")
	fmt.Scanln(&x, &y)
	// condition is used to check the point lies on x axis,y axis or at the origin
	if x == 0 && y != 0 {
		fmt.Println("The point lies on the y axis")
	} else if x != 0 && y == 0 {
		fmt.Println("The point lies on the x axis")
	} else if x == 0 && y == 0 {
		fmt.Println("The point lies at the origin")
	}
	return
}
