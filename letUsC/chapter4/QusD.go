package main

import "fmt"

/*
*This function is used to find the grace marks for a student
 */
func main() {
	var subject int
	var class string
	/* using a infinite loop loop will be end when user input will be exit */
	for count := "run"; count == "run"; {
		/* taking the user input for class of the student */
		fmt.Println("Enter the class obtained by student first, second, third or exit ")
		fmt.Scanln(&class)
		switch class {
		/* if the class is first then this case will apply*/
		case "first":
			fmt.Println("Enter the number of subjects student failed")
			fmt.Scanln(&subject)
			if subject > 3 {
				fmt.Println("Student doesn't get any grace marks")
			} else {
				fmt.Println("Student get grace of 5 marks per subject")
			}
			break
		/* if the class is second then this case will apply*/
		case "second":
			fmt.Println("Enter the number of subjects student failed")
			fmt.Scanln(&subject)
			if subject > 2 {
				fmt.Println("Student doesn't get any grace marks")
			} else {
				fmt.Println("Student get grace of 4 marks per subject")
			}
			break
		/* if the class is third then this case will apply*/
		case "third":
			fmt.Println("Enter the number of subjects student failed")
			fmt.Scanln(&subject)
			if subject > 1 {
				fmt.Println("Student doesn't get any grace marks")
			} else {
				fmt.Println("Student get grace of 5 marks per subject")
			}
			break
			/* if the class is exit then loop will be end */
		case "exit":
			count = "exit"
			break

		}
	}
	return
}
