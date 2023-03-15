package main

import (
	"fmt"
)

// In this function we are converting the distance between two cities from KiloMeter to meter,centimeter,feet and inches.
func main() {
	var distance, disMeter, disCentimeter, disFeet, disInches float64
	// we take the input from user for the distance of km
	fmt.Println("Enter the distance between two cities in KM")
	fmt.Scanln(&distance)
	// formulas for converting distance from km to centimeter, meter , feet and inches
	disMeter = distance * 1000
	disCentimeter = distance * 100000
	disFeet = distance * 3280.84
	disInches = distance * 39370.1
	// output of converted distance from Km
	fmt.Printf("Distance in meter is %f meter \n", disMeter)
	fmt.Printf("Distance in centimeter is %f centimeter \n", disCentimeter)
	fmt.Printf("Distance in feet is %f feet \n", disFeet)
	fmt.Printf("Distance in inches is %f inches \n", disInches)
	return
}
