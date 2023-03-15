package main

import "fmt"

// this function is used to check the number entered by user is odd or even
func main() {
	var num int
	// taking user input any integer to check whether it is odd or even
	fmt.Printf("Enter any integer :")
	fmt.Scanln(&num)
	// this condition use to check integer even or odd
	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
	return
}
