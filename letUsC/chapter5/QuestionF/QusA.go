package main

import "fmt"

/*
*This function is used to take user input for a float and a integer value
 */
func main() {

	/* Take user input for a float value*/
	fmt.Println("Enter a float value: ")
	var f float64
	fmt.Scanln(&f)
	/* Take user input for a integer value*/
	fmt.Println("Enter an integer value: ")
	var i int
	fmt.Scanln(&i)
	/* calling the function*/
	product := pro(f, i)
	/* output of the product of float and integer*/
	fmt.Println("The product of float and integer is ", product)
	return
}

/* This function is used to find the product of float and integer*/
func pro(f float64, i int) float64 {

	/* To find the product of float and integer*/
	product := f * float64(i)
	return product
}
