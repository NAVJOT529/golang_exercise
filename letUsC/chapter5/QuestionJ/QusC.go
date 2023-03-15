package main

import "fmt"

/*
*This function is used to count the Fibonacci sequence till 25.*/
func main() {
	num := 1
	for i := 0; i < 25; i++ {
		term := Fibonacci(num)
		num++
		fmt.Printf("%d  ", term)
	}
}

/* This function is used to create a Fibonacci sequence*/
func Fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}
