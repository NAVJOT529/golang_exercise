package main

import "fmt"

/*
*This function is used to check the factorial , prime number and odd/even of the number enter by user
 */
func main() {
	var choice string
	var number int
	/* we take a infinite loop till user input is exit when user input is exit then loop will stop */
	for count := "run"; count == "run"; {
		/* Taking the user input for what user want to solve is it factorial prime number or Odd/even or user want to exit the loop */
		fmt.Println("what do you want to check of a number, factorial, prime or odd/even choose exit if you want to exit")
		fmt.Scanln(&choice)
		/* taking a switch statement to check the number for example if user input was factorial than then program will shift for the factorial case */
		switch choice {
		/* here the the factorial case */
		case "factorial":
			factorial := 1
			fmt.Println("Enter any number")
			fmt.Scanln(&number)
			for i := 1; i <= number; i++ {
				factorial = factorial * i
			}
			fmt.Println("The factorial of the", number, "is", factorial)
			break
			/* This case is for prime numbers */
		case "prime":
			fmt.Println("Enter any number")
			fmt.Scanln(&number)
			var i int
			/* i := 2 */
			for i = 2; i <= number-1; i++ {
				if number%i == 0 {
					fmt.Println("The number", number, "is not a prime number")
					break
				}
			}
			if i == number {
				fmt.Println("The number", number, "is a prime number")
			}
			break
			/* here user can check odd or even number */
		case "odd/even":
			fmt.Println("Enter any number")
			fmt.Scanln(&number)
			if number%2 == 0 {
				fmt.Println("The number", number, "is a even number")
			} else {
				fmt.Println("The number", number, "is an odd number")
			}
			break
			/* if user input is exit then value of count will be exit and infinite loop will stop */
		case "exit":
			count = "exit"
		}
	}
	return
}
