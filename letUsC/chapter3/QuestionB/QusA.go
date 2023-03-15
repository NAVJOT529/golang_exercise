package main

import "fmt"

/*
*this function is used to calculate the overtime pay of 10 employees
 */
func main() {
	var worktime int
	/* applying loop to check and calculate overtime or overtime pay of 10 employees */
	for i := 1; i <= 10; i++ {
		fmt.Println("Enter total Work time by employee")
		fmt.Scanln(&worktime)
		if worktime > 40 {
			overtime := worktime - 40
			overtimePay := overtime * 12
			fmt.Println("your over time pay for", overtime, "hours is", overtimePay)
		} else {
			fmt.Println("Work more than 40 hours")
		}
	}
	return
}
