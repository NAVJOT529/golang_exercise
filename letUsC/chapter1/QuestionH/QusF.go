package main

import "fmt"

// this function is used to change the value of C and D which is user input
func main() {
	var C, D int
	// taking the value of C and D from user
	fmt.Println("Enter the value of C and D")
	fmt.Scanln(&C, &D)
	// changing the value of C and D
	C = C + D
	D = C - D
	C = C - D
	// output of changed value of C and D
	fmt.Printf("C=%d \nD=%d", C, D)
	return
}
