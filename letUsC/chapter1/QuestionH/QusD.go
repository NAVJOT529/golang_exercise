package main

import "fmt"

// this function is used to convert temperature from fahrenheit to degrees celsius
func main() {
	var tempFahrenheit, tempCelsius float64
	// taking the input temperature in fahrenheit from user
	fmt.Println("Enter the fahrenheit temperature of your city")
	fmt.Scanln(&tempFahrenheit)
	// calculating  temperature in celsius by using formula
	tempCelsius = (tempFahrenheit - 32) * 5 / 9
	// output for temperature in celsius
	fmt.Println("the centigrade temperature 0f your city is ", tempCelsius)
	return
}
