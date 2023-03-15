/* 12. Write a golang program to find the duplicate values of an array of integer values. */
package main

import (
	"fmt"
)

/*
*This function is used to find the duplicate values of an array
 */
func main() {
	/* array := []int{1, 2, 2, 2, 5, 8, 9, 9, 9, 9, 9, 5, 2, 2, 2, 8, 9, 6, 6}
	fmt.Println("Original values are \t", array)
	var dup []int
	numberMap := map[int]bool{}
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				numberMap[array[i]] = true
			}
		}
	}
	for i := range numberMap {
		dup = append(dup, i)
	}
	fmt.Println("duplicate arrays are\t", dup) */

	/* 	array := []int{1, 2, 2, 2, 5, 5, 5, 2, 2, 6, 8, 6}
	   	fmt.Println(array)
	   	var dup []int
	   	numberMap := map[int]bool{}
	   	for i := range array {
	   		if numberMap[array[i]] == true {
	   			dup = append(dup, array[i])
	   			numberMap[array[i]] = false
	   		} else {
	   			numberMap[array[i]] = true
	   		}
	   	}
	   	fmt.Println(dup) */
	array := []int{1, 2, 2, 2, 5, 8, 9, 9, 9, 11, 9, 9, 5, 2, 2, 2, 8, 9, 6, 6}
	fmt.Println(array)
	dup := []int{}
	numberMap := map[int]int{}
	for i := range array {
		if _, exists := numberMap[array[i]]; !exists {
			n := numberMap[array[i]]
			numberMap[array[i]] = n + 1
		} else {
			if numberMap[array[i]] == 1 {
				dup = append(dup, array[i])
			}
			numberMap[array[i]] += 1
		}
	}
	fmt.Println(numberMap)
	/* 	for i := range numberMap {
		if numberMap[i] > 1 {
			dup = append(dup, i)
		}
	} */
	fmt.Println(dup)
	return
}
