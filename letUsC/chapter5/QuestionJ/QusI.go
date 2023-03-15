package main

import (
	"fmt"
	"math"
)

/*
*This function is used to print the output (area of the triangle)
 */
func main() {
	area := CalArea()
	fmt.Println("Area of the triangle is:", area)
}

/* This function is used to calculate the area of the triangle */
func CalArea() float64 {
	var a, b, c float64
	fmt.Println("Enter the length of the sides of the triangle a,b and c")
	fmt.Scanln(&a, &b, &c)
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}
