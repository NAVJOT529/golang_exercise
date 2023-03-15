/* 31. Write a Java program to check if the sum of all the 10's in the array is exactly 30. Return false if the condition does not satisfy, otherwise true. */
package main

import "fmt"

/*
*This func is used to find the sum of all 10's value in the array , if the sum is equal to 30 then return true else false
 */
func main() {
	var arr = []int{5, 6, 7, 8, 4, 10, 4, 5, 6, 10, 7, 8, 9, 10}
	fmt.Println(arr)
	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 10 {
			sum = sum + 10
		}
	}
	fmt.Println("The sum of all 10's in array is ", sum)
	if sum == 30 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	return
}
