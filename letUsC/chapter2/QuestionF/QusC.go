package main

import "fmt"

// this function is used to check the user should be insured or not
func main() {
	var health, lives, gender string
	var age int
	// taking the user input for health , residence of gender and age
	fmt.Println("Enter your health,residence,gender and age")
	fmt.Scanln(&health, &lives, &gender, &age)
	// condition for checking the user insured or not
	if health == "excellent" && lives == "city" && gender == "male" && age >= 25 && age <= 35 {
		fmt.Println("The Premium is for Rs. 4 per thousand and Policy amount can't exceed Rs. 2 lacks ")
	} else if health == "excellent" && lives == "city" && gender == "female" && age >= 25 && age <= 35 {
		fmt.Println("The Premium is for Rs. 3 per thousand and Policy amount can't exceed Rs. 1 lacks ")
	} else if health == "poor" && lives == "village" && gender == "male" && age >= 25 && age <= 35 {
		fmt.Println("The Premium is for Rs. 6 per thousand and Policy amount can't exceed Rs. 10000")
	} else {
		fmt.Println("You are not insured")
	}
	return
}
