/* Write a golang program to find Longest Bitonic SubArray in a given array. A bitonic subArray is a subArray of a given array where elements are first sorted in increasing order, then in decreasing order. A strictly increasing or strictly decreasing subArray is also accepted as bitonic subArray.*/
package main

import "fmt"

/*
*This function is used to find the longest Bitonic subArray in a given array
 */
func main() {
	arr := []int{4, 5, 9, 5, 6, 10, 11, 9, 6, 4, 5}
	fmt.Println(arr)
	var increasing []int
	increasing = append(increasing, 1)
	for i := 1; i < len(arr); i++ {
		increasing = append(increasing, 1)
		if arr[i-1] < arr[i] {
			increasing[i] = increasing[i-1] + 1
		}
	}
	decreasing := make([]int, len(arr))
	decreasing[len(arr)-1] = 1
	for i := len(arr) - 2; i >= 0; i-- {
		decreasing[i] = 1
		if arr[i] > arr[i+1] {
			decreasing[i] = decreasing[i+1] + 1
		}
	}
	position := 1
	pointA := 0
	pointB := 0
	for i := 0; i < len(arr); i++ {
		if position < increasing[i]+decreasing[i]-1 {
			position = increasing[i] + decreasing[i] - 1
			pointA = i - increasing[i] + 1
			pointB = i + decreasing[i] - 1
		}
	}
	fmt.Println("The longest bitonic sub array is ", pointA, pointB)
	fmt.Println("The elements of the longest bitonic sub array are ", arr[pointA:pointB+1])
	fmt.Println("The length of the longest bitonic sub array is ", len(arr[pointA:pointB])+1)
	return
}
