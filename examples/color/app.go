package main

import (
	"log"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[0], 100*time.Millisecond)
	s.Prefix = "Colors: "
	if err := s.Color("yellow"); err != nil {
		log.Fatalln(err)
	}
	s.Start()
	time.Sleep(4 * time.Second)

	s.Color("red")
	s.UpdateCharSet(spinner.CharSets[20])
	s.Reverse()
	s.Restart()
	time.Sleep(4 * time.Second)

	s.Color("blue")
	s.UpdateCharSet(spinner.CharSets[3])
	s.Restart()
	time.Sleep(4 * time.Second)

	s.Color("cyan")
	s.UpdateCharSet(spinner.CharSets[28])
	s.Reverse()
	s.Restart()
	time.Sleep(4 * time.Second)

	s.Color("green")
	s.UpdateCharSet(spinner.CharSets[25])
	s.Restart()
	time.Sleep(4 * time.Second)

	s.Stop()
}
