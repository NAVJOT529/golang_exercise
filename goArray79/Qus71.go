/* 71. Write a golang program to form the largest number from a given list of non negative integers.  */
package main

import (
	"fmt"
	"math"
)

/*
*This function is used to find the largest number from given list of non negative integers
 */
func main() {
	arr := []int{1, 2, 3, 0, 4, 26}
	fmt.Println("The given array is", arr)
	fmt.Println("The largest number using said array number:")
	var array []int
	val := 0
	power := len(arr) - 1
	for permutation := make([]int, len(arr)); permutation[0] < len(permutation); nextPerm(permutation) {
		perm := append([]int{}, getPerm(arr, permutation)...)
		val = 0
		power = len(arr) - 1
		for i := 0; i < len(perm); i++ {
			val = val + int(float64(perm[i])*math.Pow(float64(10), float64(power)))
			power--
		}
		array = append(array, val)
	}
	largestNum := array[0]
	for i := range array {
		if array[i] > largestNum {
			largestNum = array[i]
		}
	}
	fmt.Println(largestNum)
	return
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
