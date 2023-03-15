/* Write a golang program to find the smallest length of a contiguous sub array of which the sum is greater than or equal to specified value. Return 0 instead */
package main

import "fmt"

/*
*This function is used to find the smallest length of a contiguous sub array of which the sum is greater than or equal to specified value.
 */
func main() {
	arr := []int{1, 2, 3, 4, 6}
	fmt.Println(arr)
	specifiedValue := 8
	minDistance := len(arr)
	distance := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] >= specifiedValue {
				distance = arr[j] - arr[i]
				if distance < minDistance {
					minDistance = distance
					if minDistance < 0 {
						minDistance = -minDistance
					}
				}
			}
		}
	}
	fmt.Println("The specified value is:", specifiedValue, "\nMinimum distance is :", minDistance)
	return
}
