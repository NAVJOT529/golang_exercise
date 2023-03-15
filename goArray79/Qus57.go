/* 57. Write a golang program to check if a sub-array is formed by consecutive integers from a given array of integers.  */
package main

import "fmt"

/*
*This function is used to check if a sub-array is formed by consecutive integers from a given array of integers
 */
func main() {
	arr := []int{2, -5, 0, 2, 9, 7, 8, 4, 3, 6, 1, -5}
	fmt.Println("The given array is:", arr)
	max := 0
	a := 0
	b := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if /* arr[i] != arr[j] && */ Check(arr[i:j+1]) == 1 {
				if len(arr[i:j+1]) >= max {
					a = i
					b = j + 1
					max = len(arr[i : j+1])
				}
			}
		}
	}
	fmt.Println("The largest sub array is", a, b-1)
	fmt.Println("The elements of the subArray is", arr[a:b])
	return
}

/* This func is used to check if their have any repeated number in the array */
func Check(arr []int) int {
	mark := 1
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				mark = 0
				break
			}
		}
	}
	return mark
}
