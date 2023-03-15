package main

import "fmt"

// this function is used to calculate the aggregate and percentage of five subjects
func main() {
	var math, chem, physics, english, computer int
	var percentage, aggregate float64
	// taking the user input for five subjects marks
	fmt.Println("Enter your marks in Math")
	fmt.Scanln(&math)
	fmt.Println("Enter your marks in Chemistry")
	fmt.Scanln(&chem)
	fmt.Println("Enter your marks in Physics")
	fmt.Scanln(&physics)
	fmt.Println("Enter your marks in English")
	fmt.Scanln(&english)
	fmt.Println("Enter your marks in Computer")
	fmt.Scanln(&computer)
	// calculating the aggregate and percentage of five subjects
	total := math + chem + physics + english + computer
	aggregate = float64(total) / 5
	percentage = float64(total) / 500 * 100
	// for output of aggregate and percentage of five subjects
	fmt.Println("The aggregate marks obtained by student is ", aggregate)
	fmt.Println("The percentage of five subjects obtained by student is ", percentage, "percent")
	return
}
