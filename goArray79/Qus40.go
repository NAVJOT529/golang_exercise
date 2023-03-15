/* 40. Write a golang program to find the two elements from a given array of positive and negative numbers such that their sum is closest to zero. */
package main

import "fmt"

/*
* This func is used to find the two elements of array their sum is closest to zero
 */
func main() {
	var arr = []int{1, 5, -6, 7, 8, -4}
	fmt.Println(arr)
	min := len(arr)
	a := 0
	b := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] < min && arr[i]+arr[j] > 0 {
				min = arr[i] + arr[j]
				a = arr[i]
				b = arr[j]
			}
		}
	}
	fmt.Println("sum of", a, "and", b, "is", min, "Which is closest to zero")
	return
}
