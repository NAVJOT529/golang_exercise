/*
 37. Write a Java program to create an array of its anti-diagonals from a given square matrix.  Go to the editor

Example:
Input :
1 2
3 4
Output:
[
[1],
[2, 3],
[4]
]
*/
package main

import "fmt"

/*
 *This function is used to create an array of its anti-diagonal
 */
func main() {
	var array = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(array)
	n := len(array)
	for i := 0; i < n-1; i++ {
		j := i
		k := 0
		for j >= 0 && k < n-1 {
			fmt.Printf("%d ", array[k][j])
			j--
			k++
		}
		fmt.Println()
	}
	for i := 0; i < n; i++ {
		a := i
		b := n - 1
		for a < n && b >= 0 {
			fmt.Printf("%d ", array[a][b])
			b--
			a++
		}
		fmt.Println()
	}
	return
}
