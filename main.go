package main

import (
	"fmt"
	"time"
)

func escrevePing() {
	for {
		fmt.Println("Ping!")
		time.Sleep(time.Millisecond * 100)
	}
}

func escrevePong() {
	for {
		fmt.Println("Pong!")
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	go escrevePing()
	go escrevePong()

	time.Sleep(time.Second * 2)
}
