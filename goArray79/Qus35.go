/* 35. Write a golang program to find the sum of the two elements of a given array which is equal to a given integer. Sample array: [1,2,4,5,6] Target value: 6 */
package main

import "fmt"

/*
*This function is used to find the pairs of elements in an array whose sum is equal to a specified number
 */
func main() {
	array := []int{1, 2, 4, 5, 6}
	fmt.Println("array:", array)
	specifiedNum := 6
	for i := 0; i < len(array)-1; i++ {
		for j := i; j < len(array)-1; j++ {
			if array[i]+array[j+1] == specifiedNum {
				fmt.Println(array[i], "and", array[j+1], "is the pair with the sum of ", specifiedNum) /* "\nIndex value of ", array[i], "is", i, "and index value of", array[j+1], "is", j+1) */
			}
		}
	}
	return
}
