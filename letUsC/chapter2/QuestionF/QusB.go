package main

import "fmt"

// this function is used to check the character enter by user is capital letter small letter digit or special symbol
func main() {
	var input rune
	var character int
	// taking the user input for enter any character
	fmt.Println("Enter any character")
	fmt.Scanf("%c", &input)
	character = int(input)
	// condition is used to check the character enter by user is capital letter small letter digit or special symbol
	if character >= 65 && character <= 90 {
		fmt.Println("You enter a capital letter")
	} else if character >= 97 && character <= 122 {
		fmt.Println("YOU enter a small letter")
	} else if character >= 48 && character <= 57 {
		fmt.Println("YOU entered a digit")
	} else if character >= 0 && character <= 47 || character >= 58 && character <= 64 || character >= 91 && character <= 96 || character >= 123 && character <= 127 {
		fmt.Println("You entered a special symbol")
	}
	return
}
