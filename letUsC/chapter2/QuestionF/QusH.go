package main

import "fmt"

// this function is used to find the efficiency of worker works in a company
func main() {
	var time int
	// taking the user input of time take by worker to complete the work
	fmt.Println("Enter the time taken by a worker to complete the work")
	fmt.Scanln(&time)
	// applying conditions to check the efficiency of the worker
	if time >= 2 && time <= 3 {
		fmt.Println("The worker is highly efficient")
	} else if time > 3 && time <= 4 {
		fmt.Println("The worker is ordered to improve the speed")
	} else if time > 4 && time <= 5 {
		fmt.Println("The worker is given training to improve his speed")
	} else if time > 5 {
		fmt.Println("The worker has to leave the Company")
	}
	return
}
