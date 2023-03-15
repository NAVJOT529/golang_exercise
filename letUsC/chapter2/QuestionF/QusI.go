package main

import "fmt"

// this function is used to check whether the student has passed, failed or is allowed to reappear in exam.
func main() {
	var marksA, marksB int
	// // taking user input of marks of subjects A and B
	fmt.Println("Enter the marks of subjects A and B")
	fmt.Scan(&marksA, &marksB)
	// applying conditions to check whether the student has passed, failed or is allowed to reappear in exam.
	if marksA >= 55 && marksB >= 45 || marksA >= 45 && marksB >= 55 {
		fmt.Println("Student has passed")
	} else if marksB < 45 && marksA >= 65 {
		fmt.Println("Student has allowed to reappear in B")
	} else {
		fmt.Println("Student has failed")
	}
	return
}
