package main

import "fmt"

/*
*This function is used to interchange the odd and even components of an array
 */
func main() {
	var array = []int{2, 9, 7, 5, 8, 4, 2, 3, 6, 5}
	fmt.Println("The original array is", array)
	for i := range array {
		if array[i]%2 != 0 {
			for j := 9; j > i; j-- {
				if array[j]%2 == 0 {
					swap(&array[i], &array[j])
					break
				}
			}
		}
	}
	fmt.Printf("After interchanging the even and odd components  ")
	for i := range array {
		fmt.Printf("%d ", array[i])
	}
}
func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
