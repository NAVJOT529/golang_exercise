package main

import "fmt"

// this function is used to calculate the total number of notes cashier have to give to withdrawer in between 100 , 50 and 10
func main() {
	var amount int
	// this user input is used to enter the amount of cash user want to withdraw
	fmt.Println("enter the amount in hundred you want to withdraw")
	fmt.Scanln(&amount)
	// to check the notes cashier have to give to user
	hundred := amount / 100
	fifty := (amount % 100) / 50
	ten := (amount % 100) % 50 / 10
	// output of no. of notes to give to the withdrawer by cashier
	fmt.Printf("you have to give %d hundreds %d fifty and %d ten notes to the withdrawer ", hundred, fifty, ten)
	return
}
