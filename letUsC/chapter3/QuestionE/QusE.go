package main

import "fmt"

/*
*This function is used to calculate the approximate level of intelligence of a person
 */
func main() {
	/*This for loop is used to produce table when y varies from 1 to 6*/
	for y := 1; y <= 6; y++ {
		/* For each value of y , x varies from 5.5 to 12.5 in steps of 0.5*/
		for x := 5.5; x <= 12.5; x = x + 0.5 {
			/* using the formula of i given in the question */
			i := 2 + (float64(y) + 0.5*float64(x))
			/* output for the values of i,x and y */
			fmt.Println("i=", i, "x=", x, "y=", y)
		}
	}
	return
}
