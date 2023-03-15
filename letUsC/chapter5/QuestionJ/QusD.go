package main

import "fmt"

/*
*This function is take the user input to find its binary equivalent
 */
func main() {
	fmt.Println("Enter the number to find its binary equivalent")
	var num int
	fmt.Scanln(&num)
	binaryNum := 0
	term := 0
	binary(num, binaryNum, term)
	return
}

/* This function is used to find the binary equivalent of the user input by recursion method */
func binary(num, binaryNum, term int) int {
	if num != 0 {
		module := num % 2
		term = term * 10
		if term == 0 {
			term = 1
		}
		binaryNum = term*module + binaryNum
		binary(num/2, binaryNum, term)

	} else {
		fmt.Println("Binary equivalent of number is :", binaryNum)
	}

	return binaryNum
}

/*
*This function is used to find the binary equivalent of the user input without recursion method
 */

/*
func binary(num, binaryNum, term int) {
	for num != 0 {
		module := num % 2
		term = term * 10
		if term == 0 {
			term = 1
		}
		binaryNum = term*module + binaryNum
		num = num / 2
	}
	fmt.Println("Binary equivalent of number is :", binaryNum)
}
*/
