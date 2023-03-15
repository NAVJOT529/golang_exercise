/* 2. Write golang program to sum values of an array. */
package main

import "fmt"

/*
*This function is used to add the values of an array
 */
func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("The original array is ", array)
	sum := 0
	for i := 0; i < len(array); i++ {
		sum = sum + array[i]
	}
	fmt.Println("The sum of the array is ", sum)
	return
}
