package main

import "fmt"

/*
*This function is use to take the input of j and k
 */
func main() {
	/* Taking the user input of j and k*/
	fmt.Println("Enter the values of j and k")
	var j, k int
	fmt.Scanln(&j, &k)
	/* calling the function*/
	CommonDiv(j, k)
	return
}

/* This function is used to find the greatest common divisor with Euclid’s algorithm*/
func CommonDiv(j, k int) {
	divide := j / k
	newK := j - divide*k
	if newK == 0 {
		fmt.Println("The greatest common divisor is :", k)
		return
	} else {
		j = k
		k = newK
		CommonDiv(j, k)
	}
}
