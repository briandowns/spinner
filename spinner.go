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

// Package spinner is a simple package to add a spinner / progress indicator to any terminal application.
package spinner

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
)

// errInvalidColor is returned when attempting to set an invalid color
var errInvalidColor = errors.New("invalid color")

// validColors holds an array of the only colors allowed
var validColors = []string{"red", "green", "yellow", "blue", "magenta", "cyan", "white"}

// validColor will make sure the given color is actually allowed
func validColor(c string) bool {
	for _, i := range validColors {
		if c == i {
			return true
		}
	}
	return false
}

// Spinner struct to hold the provided options
type Spinner struct {
	Delay      time.Duration                 // Delay is the speed of the indicator
	chars      []string                      // chars holds the chosen character set
	Prefix     string                        // Prefix is the text preppended to the indicator
	Suffix     string                        // Suffix is the text appended to the indicator
	FinalMSG   string                        // string displayed after Stop() is called
	lastOutput string                        // last character(set) written
	color      func(a ...interface{}) string // default color is white
	lock       *sync.RWMutex                 // Lock useed for
	Writer     io.Writer                     // to make testing better, exported so users have access
	active     bool                          // active holds the state of the spinner
	stopChan   chan struct{}                 // stopChan is a channel used to stop the indicator
}

// New provides a pointer to an instance of Spinner with the supplied options
func New(cs []string, d time.Duration) *Spinner {
	return &Spinner{
		Delay:    d,
		chars:    cs,
		color:    color.New(color.FgWhite).SprintFunc(),
		lock:     &sync.RWMutex{},
		Writer:   color.Output,
		active:   false,
		stopChan: make(chan struct{}, 1),
	}
}

// Start will start the indicator
func (s *Spinner) Start() {
	if s.active {
		return
	}
	s.active = true

	go func() {
		for {
			for i := 0; i < len(s.chars); i++ {
				select {
				case <-s.stopChan:
					return
				default:
					fmt.Fprint(s.Writer, fmt.Sprintf("%s%s%s ", s.Prefix, s.color(s.chars[i]), s.Suffix))
					out := fmt.Sprintf("%s%s%s ", s.Prefix, s.chars[i], s.Suffix)
					s.lastOutput = out
					s.lock.RLock()
					time.Sleep(s.Delay)
					s.lock.RUnlock()
					s.erase(out)
				}
			}
		}
	}()
}

// Stop stops the indicator
func (s *Spinner) Stop() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.active {
		s.stopChan <- struct{}{}
		s.active = false
		if s.FinalMSG != "" {
			fmt.Fprintf(s.Writer, s.FinalMSG)
		}
	}
}

// Restart will stop and start the indicator
func (s *Spinner) Restart() {
	s.Stop()
	s.Start()
}

// Reverse will reverse the order of the slice assigned to the indicator
func (s *Spinner) Reverse() {
	s.lock.Lock()
	defer s.lock.Unlock()
	for i, j := 0, len(s.chars)-1; i < j; i, j = i+1, j-1 {
		s.chars[i], s.chars[j] = s.chars[j], s.chars[i]
	}
}

// Color will set the struct field for the given color to be used
func (s *Spinner) Color(c string) error {
	if validColor(c) {
		switch c {
		case "red":
			s.color = color.New(color.FgRed).SprintFunc()
			s.Restart()
		case "yellow":
			s.color = color.New(color.FgYellow).SprintFunc()
			s.Restart()
		case "green":
			s.color = color.New(color.FgGreen).SprintFunc()
			s.Restart()
		case "magenta":
			s.color = color.New(color.FgMagenta).SprintFunc()
			s.Restart()
		case "blue":
			s.color = color.New(color.FgBlue).SprintFunc()
			s.Restart()
		case "cyan":
			s.color = color.New(color.FgCyan).SprintFunc()
			s.Restart()
		case "white":
			s.color = color.New(color.FgWhite).SprintFunc()
			s.Restart()
		default:
			return errInvalidColor
		}
	}
	return nil
}

// UpdateSpeed will set the indicator delay to the given value
func (s *Spinner) UpdateSpeed(d time.Duration) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Delay = d
}

// UpdateCharSet will change the current character set to the given one
func (s *Spinner) UpdateCharSet(cs []string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.chars = cs
}

// erase deletes written characters
func (s *Spinner) erase(a string) {
	n := utf8.RuneCountInString(a)
	s.lock.RLock()
	defer s.lock.RUnlock()
	for i := 0; i < n; i++ {
		fmt.Fprintf(s.Writer, "\b")
	}
}

// GenerateNumberSequence will generate a slice of integers at the
// provided length and convert them each to a string
func GenerateNumberSequence(length int) []string {
	var numSeq []string
	for i := 0; i < length; i++ {
		numSeq = append(numSeq, strconv.Itoa(i))
	}
	return numSeq
}
