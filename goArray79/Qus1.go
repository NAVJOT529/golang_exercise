/* 1. Write a golang program to sort a numeric array and a string array. */
package main

import (
	"fmt"
	"sort"
)

/*
*This function is used to sort a numeric array and a string array
 */
func main() {
	var numeric = []int{9, 7, 2, 8, 10, 1, 3, 4, 6, 5}
	fmt.Println("original numeric array: ", numeric)
	fmt.Printf("After numeric sort array  ")
	var pin int
	for i := 1; i < len(numeric); i++ {
		pin = numeric[i]
		j := i - 1
		for ; j >= 0 && numeric[j] > pin; j-- {
			numeric[j+1] = numeric[j]
		}
		numeric[j+1] = pin
	}
	for i := 0; i < len(numeric); i++ {
		fmt.Printf("%d ", numeric[i])
	}
	fmt.Println()
	var names = []string{"b", "e", "a", "d", "g", "c", "f"}
	fmt.Println("original string array:   ", names)
	sort.Strings(names)
	fmt.Println("After string sort array ", names)
	return
}
