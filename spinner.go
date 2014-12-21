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
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// CharSets contains the available character sets
var CharSets = [][]string{
	{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"},
	{"▖", "▘", "▝", "▗"},
	{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"},
	{"◢", "◣", "◤", "◥"},
	{"◰", "◳", "◲", "◱"},
	{"◴", "◷", "◶", "◵"},
	{"◐", "◓", "◑", "◒"},
	{".", "o", "O", "@", "*"},
	{"|", "/", "-", "\\"},
	{"◡◡", "⊙⊙", "◠◠"},
	{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
	{">))'>", " >))'>", "  >))'>", "   >))'>", "    >))'>", "   <'((<", "  <'((<", " <'((<"},
	{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
	{"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
	{"■", "□", "▪", "▫"},
	{"←", "↑", "→", "↓"},
	{"╫", "╪"},
	{"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
	{"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
	{"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
	{"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
	{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
	{"ｦ", "ｧ", "ｨ", "ｩ", "ｪ", "ｫ", "ｬ", "ｭ", "ｮ", "ｯ", "ｱ", "ｲ", "ｳ", "ｴ", "ｵ", "ｶ", "ｷ", "ｸ", "ｹ", "ｺ", "ｻ", "ｼ", "ｽ", "ｾ", "ｿ", "ﾀ", "ﾁ", "ﾂ", "ﾃ", "ﾄ", "ﾅ", "ﾆ", "ﾇ", "ﾈ", "ﾉ", "ﾊ", "ﾋ", "ﾌ", "ﾍ", "ﾎ", "ﾏ", "ﾐ", "ﾑ", "ﾒ", "ﾓ", "ﾔ", "ﾕ", "ﾖ", "ﾗ", "ﾘ", "ﾙ", "ﾚ", "ﾛ", "ﾜ", "ﾝ"},
}

// Spinner struct to hold the provided options
type Spinner struct {
	chars    []string
	Delay    time.Duration
	Prefix   string
	Suffix   string
	stopChan chan bool
	st       state
	sync.Mutex
}

type state uint8

const (
	stopped state = iota
	running
)

var ErrRunning = errors.New("spinner: already running")

var runlock sync.Mutex

// New provides a pointer to an instance of Spinner with the supplied options
func New(c []string, t time.Duration) *Spinner {
	s := &Spinner{
		Delay:    t,
		stopChan: make(chan bool, 1),
	}
	s.UpdateCharSet(c)
	return s
}

// Start will start the spinner
func (s *Spinner) Start() error {
	s.Lock()
	defer s.Unlock()
	if s.st == running {
		return ErrRunning
	}
	s.st = running
	go func() {
		runlock.Lock()
		defer runlock.Unlock()
		for {
			for i := 0; i < len(s.chars); i++ {
				select {
				case <-s.stopChan:
					return
				default:
					fmt.Printf("\r%s%s%s ", s.Prefix, s.chars[i], s.Suffix)
					time.Sleep(s.Delay)
				}
			}
		}
	}()
	return nil
}

// Stop stops the spinner
func (s *Spinner) Stop() {
	s.Lock()
	defer s.Unlock()
	if s.st == running {
		s.stopChan <- true
		s.st = stopped
	}
}

// Restart will stop and start the spinner
func (s *Spinner) Restart() {
	s.Stop()
	s.Start()
}

// Reverse will reverse the order of the slice assigned to that spinner
func (s *Spinner) Reverse() {
	s.Lock()
	defer s.Unlock()
	for i, j := 0, len(s.chars)-1; i < j; i, j = i+1, j-1 {
		s.chars[i], s.chars[j] = s.chars[j], s.chars[i]
	}
}

// UpdateSpeed is a convenience function to not have to make you
//create a new instance of the Spinner
func (s *Spinner) UpdateSpeed(delay time.Duration) { s.Delay = delay }

// UpdateCharSet will change the previously select character set to
// the provided one
func (s *Spinner) UpdateCharSet(chars []string) {
	// so that changes to the slice outside of the spinner don't change it
	// unexpectedly, create an internal copy
	s.Lock()
	defer s.Unlock()
	n := make([]string, len(chars))
	copy(n, chars)
	s.chars = n
}

// GenerateNumberSequence will generate a slice of integers at the
// provided length and convert them each to a string
func GenerateNumberSequence(length int) []string {
	numSeq := make([]string, 0)
	for i := 0; i < length; i++ {
		numSeq = append(numSeq, strconv.Itoa(i))
	}
	return numSeq
}
