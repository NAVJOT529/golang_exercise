package main

import "fmt"

// this function is used to reverse the number entered by the user
func main() {
	var num int
	// number entered by the user
	fmt.Println("Enter the five digit number")
	fmt.Scanln(&num)
	// using formula for target all digits present in the number entered by user
	num5 := num % 10
	num4 := (num / 10) % 10
	num3 := (num / 100) % 10
	num2 := (num / 1000) % 10
	num1 := (num / 10000) % 10
	// to reverse the number entered by user
	reverseNum := num5*10000 + num4*1000 + num3*100 + num2*10 + num1
	// output of reversed number entered by user
	fmt.Printf("The reverse of the number you entered is %d", reverseNum)
	return
}
