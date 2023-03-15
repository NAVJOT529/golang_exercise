package main

import "fmt"

// this function is used to check the youngest in between ram shyam and ajay
func main() {
	var ram, shyam, ajay int
	// taking user input of age of ram shyam and ajay
	fmt.Println("Enter the age of ram shyam and ajay ")
	fmt.Scanln(&ram, &shyam, &ajay)
	// this condition is used to check the yougest one
	if ram < shyam && ram < ajay {
		fmt.Println("Ram is youngest than Shyam and Ajay")
	} else if shyam < ajay && shyam < ram {
		fmt.Println("Shyam is youngest than Ram and Shyam")
	} else if ajay < ram && ajay < shyam {
		fmt.Println("Ajay is youngest than Ram and Shyam")
	}
	return
}
