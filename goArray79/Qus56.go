/* Write a golang program to sort a given binary array in linear times */
package main

import "fmt"

/*
*This func is used to sort a given binary array in linear times
 */
func main() {
	var arr = []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1}
	fmt.Println(arr)
	k := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[k] = 0
			k++
		}
	}
	for i := k; i < len(arr); i++ {
		arr[i] = 1
	}
	fmt.Println(arr)
	return
}
