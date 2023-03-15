package main

import "fmt"

/*
*This function is used to take user input for marks of three subjects and return the average and percentage
 */
func main() {
	var subject1, subject2, subject3 int
	/* Taking the user input to get the marks obtained by student in three subjects*/
	fmt.Println("Enter the marks obtained by student in three subjects")
	fmt.Scanln(&subject1, &subject2, &subject3)
	/* calling the function */
	avg := CalAverage(subject1, subject2, subject3)
	per := CalPercentage(subject1, subject2, subject3)
	/* output for the average and percentage of marks obtained by student in three subjects*/
	fmt.Println("The average marks obtained by student in three subjects is", avg)
	fmt.Println("The percentage marks obtained by student in three subjects is", per)
	return
}

/* This function is used to calculate the average */
func CalAverage(subject1, subject2, subject3 int) int {
	avg := (subject1 + subject2 + subject3) / 3
	return avg
}

/* This function is used to calculate the percentage */
func CalPercentage(subject1, subject2, subject3 int) int {
	per := ((subject1 + subject2 + subject3) / 300) * 100
	return per
}
