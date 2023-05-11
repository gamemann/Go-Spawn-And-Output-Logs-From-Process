package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var messages = []string{
	"This is a test message",
	"Another test message.",
	"Or ANOTHER test message.",
	"Or yet ANOTHER test message!!!!",
}

func Repeat() {
	for range time.Tick(time.Second * 1) {
		fmt.Println(messages[rand.Intn(len(messages))])
	}
}

func main() {
	// Repeat.
	go Repeat()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	<-sigc

	os.Exit(0)
}
