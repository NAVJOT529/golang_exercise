package main

import "fmt"

/*
*this function is used for a matchstick game being played between the computer and a user.
 */
func main() {
	var userInput int
	matchStick := 21
	fmt.Println("Lets play a matchstick game")
	fmt.Println("There are 21 matchsticks")
	fmt.Println("firstly you choose matchstick with in 1 to 4")
	fmt.Println("after that computer will choose ")
	fmt.Println("Whoever is forced to pick up the last matchstick loses the game.")

	/* creating a loop till 1 matchstick left */
	for i := 0; i <= matchStick; i++ {
		/* taking the user input for matchstick */
		fmt.Println("Pickup the matchstick within 1 to 4")
		fmt.Scanln(&userInput)
		/* computer always will choose 5-user input matchsticks*/
		computerInput := 5 - userInput
		fmt.Println("Computer choose", computerInput, "matchsticks")
		fmt.Println("Now your turn")
		matchStick = matchStick - 5
	}
	/* when one matchstick left person will loss the game who have to pickup last one. logically user will loss always */
	if matchStick == 1 {
		fmt.Println("oops,You have to choose last matchstick, so you loose the game")
		fmt.Println("better luck next time")
	}
	return
}
