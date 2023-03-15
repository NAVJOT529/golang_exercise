package main

import "fmt"

/*
*This function is used to print the multiplication table of the number entered by the user.
 */
func main() {
	var num int
	/* taking the user input for a number to find its multiplication table */
	fmt.Println("Enter the number to find its multiplication table")
	fmt.Scanln(&num)
	/* using fo loop to print the multiplication table till 10*/
	for i := 1; i <= 10; i++ {
		result := num * i
		fmt.Println(num, "*", i, "=", result)
	}
	return
}
