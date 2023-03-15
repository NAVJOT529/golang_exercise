package main

import "fmt"

/*
*This function is used to find the sum of the digits of a positive integer entered threw the keyboard recursively or non recursively
 */
func main() {
	/* Taking the user input of a positive integer*/
	fmt.Println("Enter any five digit positive integer")
	var num int
	fmt.Scanln(&num)
	/* calling the function*/
	recursive := Recursive(num)
	noRecursive := NoRecursive(num)
	/* output of the sum of the digits of the positive integer  (recursively or non recursively)*/
	fmt.Println("The sum of the digits of the positive integer (recursively)", recursive)
	fmt.Println("The sum of the digits of the positive integer (non recursively)", noRecursive)
	return
}

/* This function is used to find the sum of the digits of a positive integer by Non Recursive method*/
func NoRecursive(num int) int {
	var sum int = 0
	for num != 0 {
		module := num % 10
		sum = sum + module
		num = num / 10
	}
	return sum
}

/* This function is used to find the sum of the digits of a positive integer by Recursive method*/
func Recursive(num int) int {
	if num == 0 {
		return 0
	} else {
		module := num % 10
		sum := module + Recursive(num/10)
		return sum
	}
}
