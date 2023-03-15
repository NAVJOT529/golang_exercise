package main

import "fmt"

/*
*This function is used to copy the content of one array into another ( USED APPEND TO MAKE A ARRAY)
 */
func main() {
	/* Taking the user input for enter the number of elements in the array*/
	fmt.Println("Enter the number of elements in the array: ")
	var n int
	fmt.Scanln(&n)
	/* Taking the user input for elements of array */
	/* 	array := make([]int, n) */
	var array []int
	/* array := append(array,n) */
	for i := 0; i < n; i++ {
		var val int
		fmt.Println("Enter any element")
		fmt.Scanln(&val)
		array = append(array, val)
	}
	/* making the reverse of the elements of the array */
	/* 	reverse := make([]int, n) */
	var reverse []int
	var rev int
	for i := 0; i < n; i++ {
		/* 	reverse[(n-1)-i] = array[i] */
		rev = array[i]
		reverse = append(reverse, rev)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d", reverse[i])
	}
	return
}
