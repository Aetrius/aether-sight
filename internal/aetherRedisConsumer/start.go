package aetherRedisConsumer

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func Run() {
	fmt.Println("Running Redis Consumer")
	startConsumer()
}

var stopConsumer = make(chan struct{})

func Start() {

	// Create a ticker that ticks every 15 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Create a channel to listen for termination signals
	terminationSignal := make(chan os.Signal, 1)
	signal.Notify(terminationSignal, os.Interrupt, os.Kill)

	// Run Start function at intervals defined by the ticker
	for {
		select {
		case <-ticker.C:
			Run()
		case <-terminationSignal:
			fmt.Println("Termination signal received. Stopping image capture...")
			stopConsumer <- struct{}{} // Signal to stop image capturing
			return
		case <-stopConsumer:
			return // Stop image capture when signaled
		}
	}
}
