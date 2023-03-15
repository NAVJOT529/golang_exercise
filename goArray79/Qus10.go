/* 10. Write a golang program to find the maximum and minimum value of an array.  */
package main

import "fmt"

/*
*This function is used to find the maximum and minimum value of an array
 */
func main() {
	array := []int{5, 7, 4, 2, 3, 1, 8, 77, 55, 66, 99, 44, 22, -4, -6}
	fmt.Println(array)
	max := 0
	min := 0
	for i := range array {
		if array[i] > max {
			max = array[i]
		} else if array[i] < min {
			min = array[i]
		}
	}
	fmt.Println("The minimum value in the array is", min, "and the maximum value in the array is", max)
	return
}
