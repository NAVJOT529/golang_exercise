/* 23. Write a golang program to test the equality of two arrays. */
package main

import "fmt"

/*
*This function is used to test the equality of two arrays. Two arrays are equal if they have the same number of element
 */
func main() {
	var array1 = []int{1, 2, 3, 4, 5}
	var array2 = []int{1, 2, 3}
	fmt.Println("The first array is ", array1) /* "\nEnter the elements of second array")*/
	if len(array1) != len(array2) {
		fmt.Println("The two arrays are not equal")
		return
	}
	fmt.Println("The second array is ", array2)
	mark := 0
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			fmt.Println("The two arrays are not equal")
			mark = 2
			break
		} else if mark != 2 {
			mark = 1
		}
	}
	if mark == 1 {
		fmt.Println("The two arrays are equal")
	}
	return
}
