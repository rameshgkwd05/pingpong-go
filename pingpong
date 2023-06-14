package main

import (
	"fmt"
	"os"
	"time"
)

// The pinger prints a ping and waits for a pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		val := <-pinger
		fmt.Println("ping", val)
		time.Sleep(500 * time.Millisecond)
		ponger <- val + 1
	}
}

// The ponger prints a pong and waits for a ping
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		val := <-ponger
		fmt.Println("pong", val)
		time.Sleep(200 * time.Millisecond)
		pinger <- val + 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	// The main goroutine starts the ping/ping by sending into the ping channel
	ping <- 1

	time.Sleep(3 * time.Second)
	os.Exit(0)
}
