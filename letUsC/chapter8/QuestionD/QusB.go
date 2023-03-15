package main

import "fmt"

/*
*This function is used to find how many positive, negative, even and odd numbers present in an array
 */
func main() {
	var numbers [25]int
	/* Taking the user input for enter 25 numbers in an array */
	for i := 0; i < 25; i++ {
		fmt.Println("Enter any number")
		fmt.Scanln(&numbers[i])
	}
	negative := 0
	positive := 0
	even := 0
	odd := 0
	/* Checking how many positive,negative,even and odd numbers present in the array*/
	for i := 0; i < 25; i++ {
		if numbers[i] < 0 {
			negative++
		} else {
			positive++
		}
		if numbers[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	/* Output how many positive,negative,even and odd numbers present in the array*/
	fmt.Println("Total positive numbers Present in the array : ", positive)
	fmt.Println("Total negative numbers Present in the array : ", negative)
	fmt.Println("Total even numbers Present in the array: ", even)
	fmt.Println("Total odd numbers Present in the array: ", odd)
	return
}
