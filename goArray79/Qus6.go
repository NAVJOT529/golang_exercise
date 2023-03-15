/* 6. Write a golang program to find the index of an array element. */
package main

import "fmt"

/*
*This function is used to find the index value of a element in the array
 */
func main() {
	var array = []int{51, 52, 53, 54, 55, 56, 57, 58, 59, 60}
	fmt.Println(array)
	fmt.Println("Search any number from given array")
	var num int
	fmt.Scanln(&num)
	var index int
	for i := range array {
		if num == array[i] {
			index = i
			break
		}
	}
	fmt.Println("The index value of", num, "is", index)
	return
}
