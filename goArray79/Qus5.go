/* 5. Write a golang program to test if an array contains a specific value.  */
package main

import "fmt"

/*
*This function is used to test if an array contains a specific value
 */
func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	fmt.Println("Enter the integer you want to search in the array")
	var value int
	mark := 0
	fmt.Scanln(&value)
	for i := range array {
		if value == array[i] {
			mark = 1
			break
		}
	}
	if mark == 1 {
		fmt.Println("The number", value, "Present in the array")
	} else {
		fmt.Println("The number", value, "doesn't exist in the array")
	}
	return
}
