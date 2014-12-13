package main

// Example application using the github.com/briandowns/spinner API

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := spinner.New(spinner.CharSets[10], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(4 * time.Second)                                  // Run for some time to simulate work

	s.UpdateCharSet(spinner.CharSets[1])
	s.UpdateSpeed(200 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)

	s.UpdateCharSet(spinner.CharSets[2])
	s.UpdateSpeed(300 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)

	s.UpdateCharSet(spinner.CharSets[3])
	s.UpdateSpeed(400 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)

	s.UpdateCharSet(spinner.CharSets[4])
	s.UpdateSpeed(200 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)

	s.UpdateCharSet(spinner.CharSets[5])
	s.UpdateSpeed(100 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)
	s.Stop()
}
