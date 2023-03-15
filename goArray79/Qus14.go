/* 14. Write a golang program to find the common elements between two arrays (string values). */
package main

import "fmt"

/*
*This function is used to find the common element between two string array
 */
func main() {
	var array1 = []string{"Keyboard", "mouse", "CPU", "Vscode", "Bracket"}
	var array2 = []string{"Vscode", "Vscode", "Vscode", "Vscode", "Vscode", "Vscode", "Bracket", "PrPro", "Pycharm"}
	var dup []string
	numberMap := map[string]bool{}
	fmt.Println("Elements in array 1\t", array1)
	fmt.Println("Elements in array 2\t", array2)
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				numberMap[array1[i]] = true
			}
		}
	}
	for i := range numberMap {
		dup = append(dup, i)
	}
	fmt.Println("Duplicate values in the array 1 and array 2 of string are: ", dup)
	return
}
