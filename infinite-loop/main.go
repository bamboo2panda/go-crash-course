package main

import (
	"bufio"
	"fmt"
	"myapp/mylogger"
	"os"
	"time"
)

func main() {
	//read input from the user 5 times and write it to the log

	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go mylogger.LisstenForLog(ch)

	fmt.Println("Enter somthing")

	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
