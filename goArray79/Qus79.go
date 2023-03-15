/* 79. Write a golang program that returns the missing letter from an array of increasing letters (upper or lower). Assume there will always be one letter missing in the array. */
package main

import (
	"fmt"
)

/*
*This func is used to find the missing letter from the array
 */
func main() {
	arr := []rune{'p', 'r', 's', 't'}
	fmt.Println("The given array is:", string(arr))
	/* sort.Sort(arr) */
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] > 1 {
			fmt.Println("missing letter is:", string(arr[i]+1))
		}
	}
	return
}
