package main

import "fmt"

/*
*To take the user input into the array of 10 elements and pass the array to the modify function
 */
func main() {
	var arr [10]int
	for i := range arr {
		fmt.Println("Enter the elements of the array")
		fmt.Scanln(&arr[i])
	}
	fmt.Println("The multiplication of every element of array by 3 is", modify(arr))
	return
}

/* This function is used to modify the element of the array by multiply by 3 */
func modify(arr [10]int) [10]int {
	var multiply [10]int
	for i := range arr {
		multiply[i] = arr[i] * 3
	}
	return multiply
}
