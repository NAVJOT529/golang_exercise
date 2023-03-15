/* 3. Write a golang program to print a 10 X 10 grid. */
package main

import "fmt"

/*
*This function is used to print a 10 X 10 matrix
 */
func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("- ")
		}
		fmt.Println()
	}
	return
}
