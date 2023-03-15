package main

import "fmt"

/*
*This function is used to fill the entire screen with diamond and heart alternatively
 */
func main() {
	/* first make a infinite loop */
	for i := 0; i == i; {
		/* used the unicode for diamond ("\U0001f537") and heart ("\U0001f496")*/
		fmt.Printf("\U0001f537")
		fmt.Printf("\U0001f496")
	}
	return
}
