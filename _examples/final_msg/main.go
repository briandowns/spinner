package main

// Example application using the github.com/briandowns/spinner API

import (
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Color("red")                                              // Set the spinner color to red
	s.FinalMSG = "Complete!\nNew line!\nAnother one!\n"         // String to be displayed after Stop() is called
	s.Start()                                                   // Start the spinner
	time.Sleep(4 * time.Second)                                 // Run for some time to simulate work
	s.Stop()                                                    // Stop the spinner
}
