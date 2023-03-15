/* 43. Write a golang program to find all combination of four elements of a given array whose sum is equal to a given value. */
package main

import "fmt"

/*
*This function is used to find the combination of four elements of a given array whose sum is equal to given value
 */
func main() {
	var array = []int{5, 1, 12, 23, 44, 10, 6, 7}
	fmt.Println(array)
	sum := 24
	for i := 0; i < len(array)-3; i++ {
		for j := i + 1; j < len(array)-2; j++ {
			for k := j + 1; k < len(array)-1; k++ {
				for l := k + 1; l < len(array); l++ {
					if array[i]+array[j]+array[k]+array[l] == sum {
						fmt.Println("The combination of", array[i], array[j], array[k], "and", array[l], "is equal to", sum)
					}
				}
			}
		}
	}
	return
}
