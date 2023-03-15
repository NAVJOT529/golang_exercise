package main

import "fmt"

// this function is used to find the total illiterate men and women of a town
func main() {
	var townPopulation int = 80000
	// using formula to count total illiterate men and women
	totalLiteracy := townPopulation * 48 / 100
	totalLitMen := totalLiteracy * 35 / 100
	totalLitWomen := totalLiteracy - totalLitMen
	// for output of total illiterate men and women
	fmt.Printf("Total illiterate Men = %d\nTotal illiterate Women = %d", totalLitMen, totalLitWomen)
	return
}
