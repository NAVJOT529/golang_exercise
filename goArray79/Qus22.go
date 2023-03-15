/*22. Write a golang program to find all pairs of elements in an array whose sum is equal to a specified number */
package main

import "fmt"

/*
*This function is used to find the pairs of elements in an array whose sum is equal to a specified number */
func main() {
	array := []int{2, 7, 4, 5, -1, -5, 11, 20}
	fmt.Println(array)
	fmt.Println("Enter any specified number")
	var num int
	fmt.Scanln(&num)
	for i := 0; i < len(array)-1; i++ {
		for j := i; j < len(array)-1; j++ {
			if array[i]+array[j+1] == num {
				fmt.Println(array[i], "and", array[j+1], "is the pair with the sum of ", num)
			}
		}
	}
	return
}
