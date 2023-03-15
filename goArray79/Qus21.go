/* 21. Write a golang program to convert an ArrayList to an array.  */
package main

import "fmt"

/*
*This function is used to convert an array list to an array
 */
func main() {
	var array []int
	var val int
	var n int
	fmt.Println("Enter the size of the array")
	fmt.Scanln(&n)

	fmt.Println("Enter the elements of the array list ")
	for i := 0; i < n; i++ {
		fmt.Scanln(&val)
		array = append(array, val)
	}
	fmt.Printf("Array list entered by you : ")
	for i := range array {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()
	fmt.Println("Array :", array)
	return
}
