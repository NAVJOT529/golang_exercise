/* 4. Write a golang program to calculate the average value of array elements. */
package main

import "fmt"

/*
*This function is used to find the average of the elements of array
 */
func main() {
	var arr = [10]int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	average := sum / len(arr)
	fmt.Println("average of the elements of array is:", average)
}
