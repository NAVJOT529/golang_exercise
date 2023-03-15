package main

import "fmt"

/*
*This function is used to take user input of 25 numbers in a array . and a user input for a number to be searched and to find how many times no. present in the array
 */
func main() {
	var number [25]int
	/* Taking the user input of 25 numbers in a array*/
	fmt.Println("Enter the numbers: ")
	for i := 0; i < 25; i++ {
		fmt.Scanf("%d", &number[i])
	}
	/* Taking the user input for search a number  */
	fmt.Println("Enter the number you want to search in the array: ")
	var search int
	fmt.Scanln(&search)
	/* Counting the number of times number you searched present in the array*/
	time := 0
	for i := 0; i < 25; i++ {
		if number[i] == search {
			time = time + 1
		}
	}
	/* If time is 0 then number doesn't exist in array or if it more than 0 number exist in the array number of times*/
	if time == 0 {
		fmt.Println("Number not present in the array")
	} else {
		fmt.Println("The number", search, "present in the array", time, "times")
	}
	return
}
