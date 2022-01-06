package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const promt = "and don't type your number in, just press ENTER when ready."

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - substraction

	playTheGame(firstNumber, secondNumber, substraction, answer)
}

func playTheGame(firstNumber, secondNumber, substraction, answer int) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("Think of a number between 1 and 10", promt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, promt)
	reader.ReadString('\n')

	fmt.Println("Now multiply this result by", secondNumber, promt)
	reader.ReadString('\n')

	fmt.Println("Devide the result by the number you originaly thought of", promt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", substraction, promt)
	reader.ReadString('\n')

	fmt.Println("The answer is", answer)
}
