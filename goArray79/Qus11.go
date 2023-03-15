/* Write a golang program to reverse an array of integer values.  */
package main

import "fmt"

/*
*This function is used to reverse an array of the integer
 */
func main() {
	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("The original array is", array)
	var reverse []int
	var val int
	for i := len(array) - 1; i >= 0; i-- {
		val = array[i]
		reverse = append(reverse, val)
	}
	fmt.Println("The reversed array is", reverse)
	return
}
