/* 9. Write a golang program to insert an element (specific position) into an array. */
package main

import "fmt"

/*
* This function is used to insert an element in specific position into an array
 */
func main() {
	var array = []int{5, 9, 6, 8, 2, 3}
	fmt.Println(array)
	fmt.Println("Enter the integer you want to insert into an array")
	var num int
	fmt.Scanln(&num)
	fmt.Println("Enter the position for the element")
	var position int
	fmt.Scanln(&position)
	var newArray []int
	var value int
	for i := 0; i < len(array)+1; i++ {
		if i < position {
			value = array[i]
			newArray = append(newArray, value)
		} else if i == position {
			value = num
			newArray = append(newArray, value)
		} else {
			value = array[i-1]
			newArray = append(newArray, value)
		}
	}
	fmt.Println("The new array is ", newArray)
	return
}
