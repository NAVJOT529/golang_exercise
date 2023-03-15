/*  Write a golang program to find contiguous subArray within a given array of integers which has the largest sum */
package main

import "fmt"

/*
*This func is used to find a contiguous subArray with largest sum
 */
func main() {
	arr := []int{1, 2, -3, -4, 0, 6, 7, 8, 9}
	fmt.Println(arr)
	largestSum := 0
	for i := 0; i < len(arr)-1; i++ {
		sum := 0
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			if sum > largestSum {
				largestSum = sum
			}
		}
	}
	fmt.Println("The largest sum of contiguous sub-array: ", largestSum)
	return
}
