/* 50. Write a golang program to sort an array of positive integers of a given array, in the sorted array the value of the first element should be maximum, second value should be minimum value, third should be second maximum, fourth second be second minimum and so on. */
package main

import "fmt"

/*
*This function is used to creat a array in which the value of the first element should be maximum, second value should be minimum value, third should be second maximum, fourth second be second minimum and so on
 */
func main() {
	var arr = []int{30, 20, 40, 10, 60, 70, 50, 80, 1, 90}
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
		}
	}
	var newArray []int
	i := 0
	j := len(arr) - 1
	for i <= j {
		if j > i {
			newArray = append(newArray, arr[j])
		}
		if i < j {
			newArray = append(newArray, arr[i])
		}
		if i == j {
			newArray = append(newArray, arr[i])
		}
		j--
		i++
	}
	fmt.Println(newArray)
	return
}
