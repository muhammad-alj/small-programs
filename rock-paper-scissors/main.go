package main

import (
	"fmt"

	"github.com/muhammad-alj/rock-paper-scissors/rps"
)

func main() {
	fmt.Println("INSANE ROCK PAPER SCISSORS FIGHT HERE IT GOES!!")

	userScore, computerScore := 0, 0

	for {
		fmt.Printf("User score: %d, Computer score: %d.\n", userScore, computerScore)

		userTool := rps.PromptTool()
		computerTool := rps.RandomTool()
		winState := userTool.Fight(computerTool)

		fmt.Println()
		fmt.Printf("You chose %s, computer chose %s.\n", userTool, computerTool)

		switch winState {
		case rps.WIN:
			fmt.Println("You won!")
			userScore++
		case rps.LOSS:
			fmt.Println("You lose!")
			computerScore++
		case rps.TIE:
			fmt.Println("it's a tie!")
		}
	}
}
