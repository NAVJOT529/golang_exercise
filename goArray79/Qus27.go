/* 27. Write a golang program to find the number of even and odd integers in a given array of integers. */
package main

import "fmt"

/*
*This func is used to find the the number of even and odd integers in the array
 */
func main() {
	var arr []int
	var val int
	move := "yes"
	fmt.Println("Enter the elements of the array: ")
	for move == "yes" {
		fmt.Scanln(&val)
		arr = append(arr, val)
		fmt.Println("Do you want to enter another integer")
		fmt.Scanln(&move)
	}
	fmt.Println("Values entered by you in the array :", arr)
	even := 0
	odd := 0
	for i := range arr {
		if arr[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Println("There have ", even, "even and", odd, "odd values in the array")
	return
}
