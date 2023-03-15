package main

import "fmt"

/*
*This function is used to enter  a positive number
 */
func main() {
	var num int
	fmt.Println("Enter a positive number")
	fmt.Scanln(&num)
	Factors(num)
	return
}

/* this function is used to check the factors of the positive number entered by the user */
func Factors(num int) {
	for i := 2; num != 0; {
		if num%i == 0 {
			fmt.Printf("%d  ", i)
			num = num / i
		} else {
			i++
		}
	}
	return
}
