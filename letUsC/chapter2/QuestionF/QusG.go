package main

import (
	"fmt"
)

// this function is used to check whether the triangle is isosceles, equilateral, scalene or right angled triangle.
func main() {
	var side1, side2, side3, flag, max int
	// taking the user input for three sides of the triangle side1 side2 and side3
	fmt.Println("Enter the value of side1,side2 and side3")
	fmt.Scanln(&side1, &side2, &side3)
	// applying condition to check the triangle is isosceles, equilateral, scalene or right angled triangle.
	if side1 == side2 && side1 == side3 {
		fmt.Println("Triangle is equilateral.")
	} else if side1 == side2 || side2 == side3 || side3 == side1 {
		fmt.Println("Triangle is isosceles.")
	} else if side1 != side2 && side1 != side3 {
		fmt.Println("Triangle is scalene.")
	}
	max = side1
	if side2 > max {
		max = side2
	}
	if side3 > max {
		max = side3
	}
	if max == side1 {
		if max*max == side2*side2+side3*side3 {
			flag = 1
		}
	}
	if max == side2 {
		if max*max == side1*side1+side3*side3 {
			flag = 1
		}
	}
	if max == side3 {
		if max*max == side1*side1+side2*side2 {
			flag = 1
		}
	}
	if flag == 1 {
		fmt.Println("Triangle is right angled")
	}
	return
}
