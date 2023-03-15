package main

import "fmt"

/*
*This function is used to check the minimum life of the machine to make it a more attractive investment compared to alternative investment?
 */
func main() {
	var year int
	alternateInvestment := 120
	income := 0
	/* this loop is used to check the total year to make it a attractive investment */
	for year = 1; alternateInvestment > income; year++ {
		alternateInvestment = year * 120
		income = (1000 * year) - 4000
	}
	fmt.Println("Total year to make it a attractive investment are :", year)
	return
}
