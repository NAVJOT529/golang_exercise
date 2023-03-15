/* 58. Given two sorted arrays A and B of size p and q, write a golang program to merge elements of A with B by maintaining the sorted order i.e. fill A with first p smallest elements and fill B with remaining elements. */
package main

import "fmt"

/*
*This function is used to merge elements of A with B by maintaining the sorted order i.e. fill A with first p smallest elements and fill B with remaining elements
 */
func main() {
	A := []int{1, 3, 6, 7, 8, 10}
	B := []int{2, 4, 5, 9}
	fmt.Println(A)
	fmt.Println(B)
	lenA := len(A)
	lenB := len(B)
	var arr []int
	for i := 0; i < lenA; i++ {
		arr = append(arr, A[i])
	}
	for i := 0; i < lenB; i++ {
		arr = append(arr, B[i])
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	var newA []int
	var newB []int
	k := 0
	for i := 0; i < lenA; i++ {
		newA = append(newA, arr[k])
		k++
	}
	for i := 0; i < lenB; i++ {
		newB = append(newB, arr[k])
		k++
	}
	fmt.Println(newA)
	fmt.Println(newB)
	return
}
