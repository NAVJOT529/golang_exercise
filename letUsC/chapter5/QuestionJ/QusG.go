package main

import "fmt"

/*
*This function is used to take the user input for the value of x,y and z
 */
func main() {
	var x, y, z int
	/* taking the user input */
	fmt.Println("Enter the value of x,y and z")
	fmt.Scanln(&x, &y, &z)
	/* calling the function */
	Rotation(&x, &y, &z)
	/* Output value of x,y and z are circularly shifted  */
	fmt.Println("x=", x, "y=", y, "and z=", z)
}

/* This function is used to circularly shift the values of x,y and z */
func Rotation(x, y, z *int) {
	var a int
	a = *y
	*y = *x
	*x = *z
	*z = a
}
