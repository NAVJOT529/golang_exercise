package main

import "fmt"

/*
*This function is used to find the smallest number in the array
 */
func main() {
	var n int
	fmt.Println("Enter the size of the array")
	fmt.Scanln(&n)
	/* array := make([]int, n) */
	var array []int
	var val int
	for i := 0; i < n; i++ {
		fmt.Println("Enter the value of the array")
		fmt.Scanln(&val)
		array = append(array, val)
	}
	/* To find the smallest number in array using pointer */
	min := array[0]
	for i := range array {
		if *&array[i] <= min {
			min = *&array[i]
		}
	}
	fmt.Println("The smallest number in the array is ", min)
	return
}
