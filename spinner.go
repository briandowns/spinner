// Copyright 2014 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package spinner is a simple package to add a spinner to your application.
package spinner

import (
	"fmt"
	"strconv"
	"time"
)

// CharSets is a map of slices containing the available character sets
var CharSets = map[int][]string{
	1:  []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	2:  []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"},
	3:  []string{"▖", "▘", "▝", "▗"},
	4:  []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"},
	5:  []string{"◢", "◣", "◤", "◥"},
	6:  []string{"◰", "◳", "◲", "◱"},
	7:  []string{"◴", "◷", "◶", "◵"},
	8:  []string{"◐", "◓", "◑", "◒"},
	9:  []string{".", "o", "O", "@", "*"},
	10: []string{"|", "/", "-", "\\"},
	11: []string{"◡◡", "⊙⊙", "◠◠"},
	12: []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
	13: []string{">))'>", " >))'>", "  >))'>", "   >))'>", "    >))'>", "   <'((<", "  <'((<", " <'((<"},
	14: []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	15: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	16: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
	17: []string{"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
	18: []string{"■", "□", "▪", "▫"},
	19: []string{"←", "↑", "→", "↓"},
	20: []string{"╫", "╪"},
	21: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	22: []string{"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
	23: []string{"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
	24: []string{"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
	25: []string{"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
	26: []string{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
}

var (
	// StopChan is a bool typed channel used to stop the spinner
	StopChan = make(chan bool, 1)
)

// Spinner struct to hold the provided options
type Spinner struct {
	Chars     []string
	OrigChars []string
	Delay     time.Duration
	Offset    int
	Direction string
	Prefix,
	Suffix string
}

// New provides a pointer to an instance of Spinner with the supplied options
func New(c []string, t time.Duration) *Spinner {
	return &Spinner{
		Chars:     c,
		Delay:     t,
		Offset:    len(c) - 1,
		Direction: "right",
	}
}

// Start will start the spinner
func (s *Spinner) Start() {
	count := 0
	go func() {
		for {
			select {
			case <-StopChan:
				return
			default:
				fmt.Printf("\r%s%s%s ", s.Prefix, s.Chars[count], s.Suffix)
				time.Sleep(s.Delay)
				if count != s.Offset {
					count++
				} else {
					count = 0
				}
			}
		}
	}()
}

// Stop stops the spinner
func (s *Spinner) Stop() { StopChan <- true }

// Restart will stop and start the spinner
func (s *Spinner) Restart() {
	s.Stop()
	s.Start()
}

// Reverse will reverse the order of the slice assigned to that spinner
func (s *Spinner) Reverse() {
	var revChars []string
	s.OrigChars = s.Chars
	switch {
	case s.Direction == "right":
		for i := s.Offset; i >= 0; i-- {
			revChars = append(revChars, s.Chars[i])
		}
		s.Chars = revChars
		s.Direction = "left"
	case s.Direction == "left":
		s.Chars = s.OrigChars
		s.Direction = "right"
	}
}

// UpdateSpeed is a convenience function to not have to make you
//create a new instance of the Spinner
func (s *Spinner) UpdateSpeed(delay time.Duration) { s.Delay = delay }

// UpdateCharSet will change the previously select character set to
// the provided one
func (s *Spinner) UpdateCharSet(chars []string) { s.Chars = chars }

// GenerateNumberSequence will generate a slice of integers at the
// provided length and convert them each to a string
func GenerateNumberSequence(length int) []string {
	numSeq := make([]string, 0)
	for i := 0; i < length; i++ {
		numSeq = append(numSeq, strconv.Itoa(i))
	}
	return numSeq
}
