package main

import "fmt"

// this function is used to check the triangle is valid or not when 3 angles of triangle are user input
func main() {
	var angle1, angle2, angle3 int
	// taking user input for 3 angles of triangle and adding them
	fmt.Println("Enter the value of angle1,angle2 and angle3")
	fmt.Scanln(&angle1, &angle2, &angle3)
	sum := angle1 + angle2 + angle3
	// this condition is used to check the triangle valid or not
	if sum == 180 {
		fmt.Println("This is a valid triangle")
	} else {
		fmt.Println("this is a invalid triangle")
	}
	return
}
