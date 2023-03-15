package main

import "fmt"

/*
*This function is use to check if array is equal to its reverse
 */
func main() {
	var n int
	/* Taking the input to check the number of elements in the array */
	fmt.Println("Enter the number of elements in the array")
	fmt.Scanln(&n)
	/*
		array := make([]int, n) */
	var array []int
	var val int
	/* Taking the input of elements of the array */
	for i := 0; i < n; i++ {
		fmt.Println("Enter the element of the array")
		fmt.Scanln(&val)
		array = append(array, val)
	}

	/* To check if elements of array is equal to its reverse or not */
	mark := 0
	for i := range array {
		if array[i] != array[(n-i)-1] {
			fmt.Println("The elements of array are not equal to its reverse")
			mark = 1
			break
		} else {
			mark = 0
		}
	}
	if mark == 0 {
		fmt.Println("The elements of array is equal to its reverse")
	}
	return
}
