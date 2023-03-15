package main

import "fmt"

/*
*This function is used to generate the Prime numbers from 1 to 100
 */
func main() {
	var num [100]int
	/* putting the number 1 to 100 into a array */
	for i := 0; i < 100; i++ {
		num[i] = i + 1
	}

	/* Selecting the Prime numbers from the array */
	for i := 1; i < 100; i++ {
		if num[i] == 0 {
			continue
		}
		divide := num[i]
		for j := i + 1; j < 100; j++ {
			if num[j]%divide == 0 {
				num[j] = 0
			}
		}
	}
	/* Out put of number of prime numbers exist in the array */
	for i := 0; i < 100; i++ {
		if num[i] != 0 {
			fmt.Printf("%d ", num[i])
		}
	}
	return
}
