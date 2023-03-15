package main

import "fmt"

/*
*This function is used to take user input for a and b
 */
func main() {
	var a, b int
	/* taking the user input for value of a and b*/
	fmt.Println("Enter the value of a and b")
	fmt.Scanln(&a, &b)
	/* calling the function*/
	value := Power(a, b)
	fmt.Println("The value of", a, "raised to power", b, "is:", value)
	return
}

/* This function is used to calculate the value of a raised to power b */
func Power(a int, b int) int {
	value := 1
	for i := 1; i <= b; i++ {
		value = value * a
	}
	return value
}
