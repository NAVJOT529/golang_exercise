/* 69. Write a golang program to find minimum sub array sum of specified size in a given array of integers. */
package main

import "fmt"

/*
*This function is used to find minimum sub array sum of specified size in a given array of integers
 */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("The given array is", arr)
	size := 4
	minimumSum := len(arr)
	a := 0
	b := 0
	for i := 0; i <= len(arr)-size; i++ {
		sum := 0
		for j := i; j < i+size; j++ {
			sum = sum + arr[j]
		}
		if sum <= minimumSum {
			minimumSum = sum
			a = i
			b = i + (size - 1)
		}
	}
	fmt.Println("The size of the array is", size, "\nSub array from", a, "to", b, "has the minimum sum", minimumSum)
	return
}
