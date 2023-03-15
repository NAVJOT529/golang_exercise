/* 7. Write a golang program to remove a specific element from an array. */
package main

import "fmt"

/*
*This function is used to remove a specific element from an array
 */
func main() {
	array := []int{40, 37, 97, 64, 54, 12, 50, 20, 27}
	fmt.Println(array)
	fmt.Println("Enter the number you want to remove from given array")
	var num int
	fmt.Scanln(&num)
	for i := 0; i < len(array); {
		if num != array[i] {
			i++
		}
		if num == array[i] {
			break
		}
		if i == len(array)-1 {
			fmt.Println("number doesn't exit in array")
			return
		}
	}
	var mark int
	var index int
	for i := range array {
		if num != array[i] {
			continue
		}
		if num == array[i] {
			index = i
			mark = 0
		}
	}
	if mark != 1 {
		newLength := 0
		for i := range array {
			if index != i {
				array[newLength] = array[i]
				newLength++
			}
		}
		newArray := array[:newLength]
		fmt.Println(newArray)
	}
	return
}
