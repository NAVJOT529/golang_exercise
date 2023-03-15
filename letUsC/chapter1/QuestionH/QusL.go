package main

import "fmt"

// this function is used to find the cost price of one item
func main() {
	var sellingPrice, profit int
	// taking the input from user for selling price of 15 items and profit earned on them
	fmt.Println("Enter the selling price of 15 items and profit earned on them")
	fmt.Scanln(&sellingPrice, &profit)
	// finding the cost price of one item
	costPrice := sellingPrice - profit
	oneItemCostPrice := costPrice / 15
	// output of cost price of single item
	fmt.Println("The cost price of one item is ", oneItemCostPrice)
	return
}
