package main

import (
	"fmt"
	"math"
)

// this function is used to check the  a point lies on the circle inside the circle or outside the circle
func main() {
	var x1, x2, y1, y2, radius float64
	// taking the user input of the coordinates of center of circle and the point and radius of the circle
	fmt.Println("Enter the coordinates of the center of the circle x1,y1")
	fmt.Scanln(&x1, &y1)
	fmt.Println("Enter the coordinates of the point x2,y2")
	fmt.Scanln(&x2, &y2)
	fmt.Println("Enter the radius of the circle")
	fmt.Scanln(&radius)
	// calculating the distance from center of circle to the point
	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	// condition to check the point lies inside, outside or on the circle
	if distance < radius {
		fmt.Println("The point lies inside the circle")
	} else if distance > radius {
		fmt.Println("The point lies outside the circle")
	} else {
		fmt.Println("The point lies on the circle")
	}
	return
}
