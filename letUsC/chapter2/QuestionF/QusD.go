package main

import "fmt"

// this function is use to check the grade of the steel
func main() {
	var hardness, tensileStrength int
	var carbonContent float64
	// taking the user input for hardness, tensile strength and carbon content of steel
	fmt.Println("Enter the hardness, tensile strength and carbon content of steel ")
	fmt.Scanln(&hardness, &tensileStrength, &carbonContent)
	// conditions to check the grade of the steel
	if hardness > 50 && carbonContent < 0.7 && tensileStrength > 5600 {
		fmt.Println("The grade of the steel is 10")
	} else if hardness > 50 && carbonContent < 0.7 && tensileStrength < 5600 {
		fmt.Println("The grade of the steel is 9")
	} else if hardness < 50 && carbonContent < 0.7 && tensileStrength > 5600 {
		fmt.Println("The grade of the steel is 8")
	} else if hardness > 50 && carbonContent > 0.7 && tensileStrength > 5600 {
		fmt.Println("The grade of the steel is 7")
	} else if hardness > 50 && carbonContent > 0.7 && tensileStrength < 5600 || hardness < 50 && carbonContent > 0.7 && tensileStrength > 5600 || hardness < 50 && carbonContent < 0.7 && tensileStrength < 5600 {
		fmt.Println("The grade of the steel is 6")
	} else {
		fmt.Println("The grade of the steel is 5")
	}
	return
}
