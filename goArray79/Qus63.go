/* 63. Write a golang program to replace each element of the array with product of every other element in a given array of integers.  */
package main

import "fmt"

/*
*This func is used to replace the each element of the array with the product of every other element in a given array of integers
 */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	product := 1
	var array []int
	for i := range arr {
		for j := range arr {
			if i == j {
				continue
			} else {
				product = product * arr[j]
			}
		}
		array = append(array, product)
		product = 1
	}
	fmt.Println(array)
	return
}
