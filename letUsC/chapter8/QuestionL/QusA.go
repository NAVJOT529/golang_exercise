package main

import "fmt"

/*
*This function is used to initialize a three dimensional array threed[3][2][3] and refer the first and last element in the array
 */
func main() {
	var threed = [3][2][3]int{{{24, 0, 0}, {0, 0, 0}}, {{0, 0, 0}, {0, 0, 0}}, {{0, 0, 0}, {0, 0, 94}}}
	fmt.Println("The first element of the array is", threed[0][0][0])
	fmt.Println("The last element of the array is", threed[2][1][2])
	return
}
