/* 13. Write a golang program to find the duplicate values of an array of string values. */
package main

import "fmt"

/*
*This function is used to find the duplicate strings in a array
 */
func main() {
	var array = []string{"India", "Canada", "Canada", "Canada", "Canada", "India", "India", "India", "India", "Dubai", "Canada", "America", "India"}
	fmt.Println("This is the list of countries:\t", array)
	/* var dup []string
	stringMap := map[string]bool{}
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				stringMap[array[i]] = true
			}
		}
	}
	for i := range stringMap {
		dup = append(dup, i)
	}
	fmt.Println("Duplicate values in the array of string are: ", dup) */
	dup := []string{}
	stringMap := map[string]int{}
	for i := range array {
		if _, exists := stringMap[array[i]]; exists {
			n := stringMap[array[i]]
			stringMap[array[i]] = n + 1
		} else {
			stringMap[array[i]] = 1
		}

	}
	fmt.Println(stringMap)
	for i := range stringMap {
		if stringMap[i] > 1 {
			dup = append(dup, i)
		}
	}
	fmt.Println(dup)
	return
}
