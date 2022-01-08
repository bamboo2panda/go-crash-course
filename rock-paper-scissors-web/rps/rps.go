package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
	// PLAYERWINS   = 1
	// COMPUTERWINS = 2
	// DRAW         = 3
)

type Round struct {
	// Winner         int    `json:"winner"`
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

var winMessages = []string{
	"Good job!",
	"Nice work!",
	"You should buy a lottery ticket!",
}

var loseMessages = []string{
	"Too bad!",
	"Try again!",
	"This is just not your day.",
}

var drawMessages = []string{
	"Great minds think alike.",
	"Uh oh. Try again.",
	"Nobody wins, but you can try again.",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())

	computerValue := rand.Intn(3)
	computerChoise := ""
	roundResult := ""
	// winner := 0

	switch computerValue {
	case ROCK:
		computerChoise = "Computer chose ROCK"
		break
	case PAPER:
		computerChoise = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoise = "Computer chose SCISSORS"
		break
	default:
	}

	messageInt := rand.Intn(3)
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		message = winMessages[messageInt]
	} else {
		roundResult = "Computer win!"
		message = loseMessages[messageInt]
	}

	var result Round
	result.Message = message
	result.ComputerChoice = computerChoise
	result.RoundResult = roundResult

	return result
}
