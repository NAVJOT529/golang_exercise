/* 32. Write a golang program to check if an array of integers contains two specified elements 65 and 77.  */
package main

import "fmt"

/*
*This function is used to check if an array contains two specified elements 65 and 77
 */
func main() {
	var arr = []int{5, 9, 7, 6, 4, 65, 7, 77}

	fmt.Println(arr)
	sixtyFive := 0
	seventySeven := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 65 {
			sixtyFive = 1
		}
		if arr[i] == 77 {
			seventySeven = 1
		}
	}
	if sixtyFive == 1 && seventySeven == 1 {
		fmt.Println("yes ,Array contains two specified elements")
	} else {
		fmt.Println("Array doesn't contain two specified elements")
	}
	return
}
