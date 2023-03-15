package main

import "fmt"

// this function is used to check the loss or profit get by user when selling price and cost price is input threw keyboard
func main() {
	var sellingPrice, costPrice int
	// take the user input for cost price and selling price of the item
	fmt.Println("Enter the cost price and selling price of the item")
	fmt.Scanln(&costPrice, &sellingPrice)
	// this condition for checking the profit or loss get by user
	if sellingPrice > costPrice {
		profit := sellingPrice - costPrice
		fmt.Printf("you are in profit with %d rupees", profit)
	} else {
		loss := costPrice - sellingPrice
		fmt.Printf("you are in loss with %d rupees", loss)
	}
	return
}
