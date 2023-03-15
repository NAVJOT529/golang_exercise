/* 72. Write a golang program to find and print one continuous sub Array (from a given array of integers) that if you only sort the said sub Array in ascending order then the entire array will be sorted in ascending order.  */
package main

import "fmt"

/*
*This function is used to find the sub array which need to be sorted in ascending order then the entire array will be sorted in ascending order
 */
func main() {
	arr := []int{1, 2, 3, 0, 4, 6}
	fmt.Println(arr)
	a := 0
	for a = 0; a < len(arr)-1; a++ {
		if arr[a] > arr[a+1] {
			break
		}
	}
	b := 0
	for b = len(arr) - 1; b > 0; b-- {
		if arr[b] < arr[b-1] {
			break
		}
	}
	max := arr[a]
	min := arr[a]
	for i := a + 1; i <= b; i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	for i := 0; i < a; i++ {
		if arr[i] > min {
			a = i
			break
		}
	}
	for i := len(arr) - 1; i > b+1; i-- {
		if arr[i] < max {
			b = i
			break
		}
	}
	fmt.Println(arr[a : b+1])
	return
}
