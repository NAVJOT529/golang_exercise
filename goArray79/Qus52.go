/* 52. Write a golang program to separate even and odd numbers of a given array of integers. Put all even numbers first, and then odd numbers */
package main

import "fmt"

/*
*This function is used to separate even and odd numbers of the array
 */
func main() {
	var arr = []int{2, 3, 5, 4, 6, 8, 10, 9, 7, 5}
	fmt.Println("The give array is ", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]%2 != 0 {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println("The new array is :", arr)
	return
}
