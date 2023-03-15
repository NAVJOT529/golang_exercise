package main

import "fmt"

/*
* this function is used to print all the Ascii values
 */
func main() {
	var character rune
	for character = 0; character < 255; character++ {
		fmt.Println("Ascii value of", string(character), "is", character)
	}
	return
}
