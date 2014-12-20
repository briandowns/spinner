package main

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	fmt.Println("I'm about to do work")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "Doing the work... "
	s.Start()
	time.Sleep(2 * time.Second)
	s.Stop()
	fmt.Println("Done!")
}
