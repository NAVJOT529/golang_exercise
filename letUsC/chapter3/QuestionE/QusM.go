package main

import (
	"fmt"
	"math"
)

/*
* This function is used to calculate the sum of first seven terms
 */
func main() {
	var x float64
	finalResult := 0.0
	/* firstly taking the user input to take the value of x*/
	fmt.Println("Enter value of x")
	fmt.Scanln(&x)
	/* calculating the first term now six terms are pending to calculate*/
	firstTerm := x - 1/x
	/* this loop is used to calculate next 6 terms in the series */
	for term := 2; term <= 7; term++ {
		power := math.Pow(firstTerm, float64(term))
		termResult := power / 2
		finalResult = finalResult + termResult
	}
	/*finalResult is the sum of all the seven terms in the series*/
	finalResult = firstTerm + finalResult
	fmt.Println("The sum of first seven terms is ", finalResult)
	return
}
