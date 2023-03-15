package main

import "fmt"

/*
*This function is used to enter  a positive number
 */
func main() {
	var num int
	fmt.Println("Enter a positive number")
	fmt.Scanln(&num)
	i := 2
	Factors(num, i)
	return
}

/* this function is used to check the Prime factors of the positive number entered by the user recursively */
func Factors(num, i int) {
	if num == 0 {
		return
	} else if num%i == 0 {
		fmt.Printf("%d ", i)
		Factors(num/i, i)
	} else {
		Factors(num, i+1)
	}
	return
}
