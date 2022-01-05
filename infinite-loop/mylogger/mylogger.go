package mylogger

import "log"

func LisstenForLog(ch chan string) {
	for {
		msg := <-ch
		log.Println(msg)
	}
}
