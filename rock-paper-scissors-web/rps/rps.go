package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

func PlayRound(playerValue int) (int, string, string) {
	rand.Seed(time.Now().UnixNano())

	computerValue := rand.Intn(3)
	computerChoise := ""
	roundResult := ""
	winner := 0

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

	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer win!"
		winner = COMPUTERWINS
	}

	return winner, computerChoise, roundResult
}