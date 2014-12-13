package main

// Example application using the github.com/briandowns/spinner API

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := New(spinner.CharSets[10], 100*time.Millisecond) // Build our new spinner
	s.Start()                                            // Start the spinner
	time.Sleep(5 * time.Second)                          // Run for some time to simulate work
	s.Stop()                                             // Stop the spinner
}
