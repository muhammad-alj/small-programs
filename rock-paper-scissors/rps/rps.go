package rps

import (
	"fmt"
	"math/rand"
)

// tool
type tool string

const (
	rock     tool = "rock"
	paper    tool = "paper"
	SCISSORS tool = "scissors"
)

var tools = []tool{rock, paper, SCISSORS}

func (t tool) Fight(other tool) winState {
	if t == other {
		return TIE
	}

	isWinner := (t == rock && other == SCISSORS) ||
		(t == SCISSORS && other == paper) ||
		(t == paper && other == rock)

	if isWinner {
		return WIN
	}
	return LOSS
}

func PromptTool() tool {
	toolsLength := len(tools)

	fmt.Println("Choose tool.")
	for i := 0; i < toolsLength; i++ {
		fmt.Printf("%d) %s\n", i+1, tools[i])
	}

	// The input should be between 1-toolsLength for better UX
	var input int

	for {
		fmt.Print("> ")

		if _, err := fmt.Scanln(&input); err != nil {
			// this clears the input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if input <= toolsLength && input >= 1 {
			break
		}
	}

	// subtract one so the input becomes a proper index
	return tools[input-1]
}

func RandomTool() tool {
	return tools[rand.Intn(len(tools))]
}

// Win State
type winState int

const (
	WIN winState = iota
	LOSS
	TIE
)
