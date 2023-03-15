package main

import (
	"fmt"
	"math"
)

/*
* This function is used to taking the input of the coordinates of three points of a triangle
* and check the point (x,y) lies inside the triangle or not
 */
func main() {
	var x1, y1, x2, y2, x3, y3 float64
	/* taking the user input for coordinates of three point of a triangle*/
	fmt.Println("Enter the coordinates of three points of a triangle")
	fmt.Println("Enter the values coordinates x1,y1")
	fmt.Scanln(&x1, &y1)
	fmt.Println("Enter the values coordinates x2,y2")
	fmt.Scanln(&x2, &y2)
	fmt.Println("Enter the values coordinates x3,y3")
	fmt.Scanln(&x3, &y3)
	/* distance of the three points of a triangle */
	A := Distance(x1, y1, x2, y2)
	B := Distance(x2, y2, x3, y3)
	C := Distance(x1, y1, x3, y3)
	/* for find the area of the triangle*/
	area := CalArea(A, B, C)
	/* output of the area of the triangle*/
	fmt.Println("The area of the triangle is : ", area)

	/* taking the input for a point  and check is this point lies inside the triangle*/
	fmt.Println("Enter the the coordinates of a point x,y")
	var x, y float64
	fmt.Scanln(&x, &y)
	/* distance of all points of the triangle and given point*/
	a := Distance(x1, y1, x, y)
	b := Distance(x2, y2, x, y)
	c := Distance(x3, y3, x, y)
	/* area of the triangles made with the give point and points of the triangle */
	area1 := CalArea(A, b, c)
	area2 := CalArea(B, b, c)
	area3 := CalArea(c, a, c)
	/* For check is given point lies inside point or not */
	check := Cheackline(area, area1, area2, area3)
	if check == 1 {
		fmt.Println("The point xy lies inside the triangle")
	} else {
		fmt.Println("The point xy doesn't lies inside the triangle")
	}

}

/* This function is used to find the distance*/
func Distance(x1, y1, x2, y2 float64) float64 {
	dis := math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
	return dis
}

/* This function is used to find the area of the triangle*/
func CalArea(a, b, c float64) float64 {
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

/* This function is used to check if the point lies inside the triangle or not*/
func Cheackline(a, b, c, d float64) int {
	var mark int
	if a == b+c+d {
		mark = 1
	} else {
		mark = 0
	}
	return mark
}
