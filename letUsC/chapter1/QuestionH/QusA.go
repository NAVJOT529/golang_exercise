package main

import (
	"fmt"
)

/*
* In this function we are count the grossSalary of Ramesh
 */
func main() {
	var basicSalary, grossSalary float64
	fmt.Printf("Enter Ramesh's basic salary \n")
	fmt.Scanln(&basicSalary)
	/* here  we are adding the house rent allowance and dearness allowance of ramesh*/
	grossSalary = 0.4*basicSalary + 0.2*basicSalary + basicSalary
	fmt.Println("The gross salary of Ramesh is: ", grossSalary)
	return
}
