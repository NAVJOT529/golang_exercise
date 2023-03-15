package main

import "fmt"

// this function is used to calculate the area and perimeter of the rectangle and And area and circumference of circle
func main() {
	var length, breadth, radius float64
	// taking the user input for length and breadth of rectangle and radius of the circle
	fmt.Println("Enter the length of the rectangle")
	fmt.Scanln(&length)
	fmt.Println("Enter the breadth of the rectangle")
	fmt.Scanln(&breadth)
	fmt.Println("Enter the radius  of the circle")
	fmt.Scanln(&radius)
	// output for area and perimeter of the rectangle And area and circumference of the circle
	fmt.Printf("Area of the rectangle is %f\n", length*breadth)
	fmt.Printf("Perimeter of the rectangle is %f\n", 2*(length+breadth))
	fmt.Printf("Circumference of the circle is %f\n", 2*3.14*radius)
	fmt.Printf("Area of the circle is %f\n", 3.14*radius*radius)
	return
}
