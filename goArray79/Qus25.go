/* 25. Write a golang program to find common elements from three sorted (in non-decreasing order) arrays. */
package main

import (
	"fmt"
	"sort"
)

/*
*This function is used to find the common elements from three sorted array
 */
func main() {
	var array1 = []int{2, 1, 8, 4}
	var array2 = []int{2, 3, 4, 8, 10, 16}
	var array3 = []int{4, 8, 14, 40}
	fmt.Println("array1:\t\t", array1)
	fmt.Println("array2:\t\t", array2)
	fmt.Println("array3:\t\t", array3)
	sort.Ints(array1)
	sort.Ints(array2)
	sort.Ints(array3)
	var common []int
	/* tanking x for element of first array y for elements of second array and z for elements of third array */
	x := 0
	y := 0
	z := 0
	for x < len(array1) && y < len(array2) && z < len(array3) {
		if array1[x] == array2[y] && array2[y] == array3[z] {
			common = append(common, array1[x])
			x++
			y++
			z++
		} else if array1[x] < array2[y] {
			x++
		} else if array2[y] < array3[z] {
			y++
		} else {
			z++
		}
	}
	fmt.Println("common elements: ", common)
	return
}

/* 	var dup []int
   	numberMap := map[int]bool{}
   	for i := range array1 {
   		if numberMap[array1[i]] == true {
   			dup = append(dup, array1[i])
   			numberMap[array1[i]] = false
   		} else {
   			numberMap[array1[i]] = true
   		}
   	}
   	for i := range array2 {
   		if numberMap[array2[i]] == true {
   			dup = append(dup, array2[i])
   			numberMap[array2[i]] = false
   		} else {
   			numberMap[array2[i]] = true
   		}
   	}
   	for i := range array3 {
   		if numberMap[array3[i]] == true {
   			dup = append(dup, array3[i])
   			numberMap[array3[i]] = false
   		} else {
   			numberMap[array3[i]] = true
   		}
   	}
   	fmt.Println(dup) */
