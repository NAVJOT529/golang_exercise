/* 39. Write a golang program to print all the LEADERS in the array. Note: An element is leader if it is greater than all the elements to its right side. */
package main

import "fmt"

/*
*This function is used to Print all the leaders in the array
 */
func main() {
	var arr = []int{10, 9, 14, 23, 15, 0, 9}
	fmt.Println(arr)
	var leaders []int
	mark := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				mark = 1
			} else {
				mark = 0
				break
			}
		}
		if mark == 1 {
			leaders = append(leaders, arr[i])
		}
	}
	leaders = append(leaders, arr[len(arr)-1])
	fmt.Println("All the leaders in the array: ", leaders)
	return
}
