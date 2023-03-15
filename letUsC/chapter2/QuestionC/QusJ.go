package main

import "fmt"

// this function is used to check the three points lies on straight line or not
func main() {
	var x1, x2, x3, y1, y2, y3 int
	// taking the user input of coordinates of three points
	fmt.Println("enter the value of coordinates of first point x1,y1")
	fmt.Scanln(&x1, &y1)
	fmt.Println("enter the value of coordinates of second point x2,y2")
	fmt.Scanln(&x2, &y2)
	fmt.Println("enter the value of coordinates of first point x3,y3")
	fmt.Scanln(&x3, &y3)
	// this condition is used to check three points lies on straight line or not
	if (y2-y1)/(x2-x1) == (y3-y1)/(x3-x1) {
		fmt.Println("All the three points lies on straight line")
	} else {
		fmt.Println("All the three points are not lies on straight line")
	}
	return
}
