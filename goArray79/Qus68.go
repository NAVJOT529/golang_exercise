/* Write a golang program to create all possible permutations of a given array of distinct integers. */
package main

import "fmt"

/*
*This function is used to print all the permutations of a given array
 */

func main() {
	arr := []int{1, 2, 3, 4}
	for permutation := make([]int, len(arr)); permutation[0] < len(permutation); nextPerm(permutation) {
		result := append([]int{}, getPerm(arr, permutation)...)
		fmt.Println(result)
	}
}

/* This function is used to find the next permutation */
func nextPerm(permutation []int) {
	for i := len(permutation) - 1; i >= 0; i-- {
		if i == 0 || permutation[i] < len(permutation)-i-1 {
			permutation[i]++
			return
		}
		permutation[i] = 0
	}
}

/* This function is used to get the next permutation */
func getPerm(arr, permutation []int) []int {
	for i, v := range permutation {
		arr[i], arr[i+v] = arr[i+v], arr[i]
	}
	return arr
}
