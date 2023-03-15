/* 20. Write a golang program to convert an array to ArrayList. */
package main

import "fmt"

/*
*This function is used to a convert a array into array list
 */
func main() {
	var array = []int{20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40}
	fmt.Println("Array :\t", array)
	fmt.Print("This is the array list of array :")
	for i := range array {
		fmt.Printf("%d ", array[i])
	}
	return
}
