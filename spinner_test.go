// Copyright (c) 2021 Brian J. Downs
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

package spinner

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/term"
)

const baseWait = 3

// syncBuffer
type syncBuffer struct {
	sync.Mutex
	bytes.Buffer
}

// Write
func (b *syncBuffer) Write(data []byte) (int, error) {
	b.Lock()
	defer b.Unlock()
	return b.Buffer.Write(data)
}

// withOutput
func withOutput(a []string, d time.Duration) (*Spinner, *syncBuffer) {
	var out syncBuffer
	s := New(a, d)
	s.Writer = &out
	return s, &out
}

// TestNew verifies that the returned instance is of the proper type
func TestNew(t *testing.T) {
	for i := 0; i < len(CharSets); i++ {
		s := New(CharSets[i], 1*time.Second)
		if reflect.TypeOf(s).String() != "*spinner.Spinner" {
			t.Errorf("New returned incorrect type kind=%d", i)
		}
	}
}

// TestStart will verify a spinner can be started
func TestStart(t *testing.T) {
	s := New(CharSets[1], 100*time.Millisecond)
	s.Color("red")
	s.Start()
	time.Sleep(baseWait * time.Second)
	s.Stop()
	time.Sleep(100 * time.Millisecond)
}

// TestActive will verify we can tell when a spinner is running
func TestActive(t *testing.T) {
	if fd := os.Stdout.Fd(); !term.IsTerminal(int(fd)) {
		t.Log("not running in a terminal")
		return
	}
	s := New(CharSets[1], 100*time.Millisecond)
	if s.Active() {
		t.Error("expected a new spinner to not be active")
	}
	s.Start()
	if !s.Active() {
		t.Error("expected a started spinner to be active")
	}
	s.Stop()
	if s.Active() {
		t.Error("expected a stopped spinner to be active")
	}
}

// TestStop will verify a spinner can be stopped
func TestStop(t *testing.T) {
	p, out := withOutput(CharSets[14], 100*time.Millisecond)
	p.Color("yellow")
	p.Start()
	time.Sleep(500 * time.Millisecond)
	p.Stop()
	// because the spinner will print an appropriate number of backspaces before stopping,
	// let it complete that sleep
	time.Sleep(100 * time.Millisecond)
	out.Lock()
	len1 := out.Len()
	out.Unlock()
	time.Sleep(300 * time.Millisecond)
	out.Lock()
	defer out.Unlock()
	len2 := out.Len()
	if len1 != len2 {
		t.Errorf("expected equal, got %v != %v", len1, len2)
	}
	p = nil
}

// TestRestart will verify a spinner can be stopped and started again
func TestRestart(t *testing.T) {
	s, out := withOutput(CharSets[4], 40*time.Millisecond)

	s.Start()
	time.Sleep(150 * time.Millisecond)
	s.Restart()
	time.Sleep(158 * time.Millisecond)
	s.Stop()
	time.Sleep(10 * time.Millisecond)

	result := out.Bytes()
	first := result[:len(result)/2]
	second := result[len(result)/2:]
	if !bytes.Equal(first, second) {
		t.Errorf("expected restart output to match initial output. got=%q want=%q", first, second)
	}
}

// TestRapidStopRestart tests that calling Stop and Restart rapidly in sequence
// and concurrently does not deadlock or block subsequent Start/Restart goroutines.
func TestRapidStopRestart(t *testing.T) {
	s := New(CharSets[1], 2*time.Millisecond)
	s.Writer = ioutil.Discard
	if f, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0); err == nil {
		defer f.Close()
		s.WriterFile = f
	}

	var updates int32
	s.PreUpdate = func(_ *Spinner) {
		atomic.AddInt32(&updates, 1)
	}

	// Rapid sequential Stop and Restart (Issue #101 regression)
	s.Start()
	for i := 0; i < 50; i++ {
		s.Stop()
		s.Restart()
		time.Sleep(5 * time.Millisecond)
	}
	s.Stop()

	if term.IsTerminal(int(s.WriterFile.Fd())) && atomic.LoadInt32(&updates) == 0 {
		t.Error("expected spinner to have performed updates across rapid restarts")
	}
	t.Logf("Total updates during rapid restarts: %d", atomic.LoadInt32(&updates))

	// Rapid concurrent Stop and Restart under race detector
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				s.Restart()
				time.Sleep(time.Millisecond)
				s.Stop()
			}
		}()
	}
	wg.Wait()
}

func TestDisable(t *testing.T) {
	s, _ := withOutput(CharSets[4], 100*time.Millisecond)

	s.Start()
	time.Sleep(150 * time.Millisecond)
	if !s.Enabled() {
		t.Error("expected enabled spinner after startup")
	}
	time.Sleep(150 * time.Millisecond)
	s.Disable()
	time.Sleep(150 * time.Millisecond)
	if s.Enabled() {
		t.Error("expected disabling the spinner works")
	}
	time.Sleep(150 * time.Millisecond)
	s.Enable()
	time.Sleep(150 * time.Millisecond)
	if !s.Enabled() {
		t.Error("expected enabling the spinner works")
	}
}

// TestHookFunctions will verify that hook functions works as expected
func TestHookFunctions(t *testing.T) {
	if fd := os.Stdout.Fd(); !term.IsTerminal(int(fd)) {
		t.Log("not running in a terminal")
		return
	}
	s := New(CharSets[4], 50*time.Millisecond)
	var out syncBuffer
	s.Writer = &out
	s.PreUpdate = func(s *Spinner) {
		fmt.Fprintf(s.Writer, "pre-update")
	}
	s.PostUpdate = func(s *Spinner) {
		fmt.Fprintf(s.Writer, "post-update")
	}

	s.Start()
	s.Color("cyan")
	time.Sleep(200 * time.Millisecond)
	s.Stop()
	time.Sleep(50 * time.Millisecond)
	out.Lock()
	defer out.Unlock()
	result := out.Bytes()
	if !bytes.Contains(result, []byte("pre-update")) {
		t.Error("pre-update failed")
	}

	if !bytes.Contains(result, []byte("post-update")) {
		t.Error("post-update failed")
	}
	s = nil
}

// TestReverse will verify that the given spinner can stop and start again reversed
func TestReverse(t *testing.T) {
	a := New(CharSets[10], 1*time.Second)
	a.Color("red")
	a.Start()
	time.Sleep(baseWait * time.Second)
	a.Reverse()
	a.Restart()
	time.Sleep(baseWait * time.Second)
	a.Reverse()
	a.Restart()
	time.Sleep(baseWait * time.Second)
	a.Stop()
	a = nil
}

// TestUpdateSpeed verifies that the delay can be updated
func TestUpdateSpeed(t *testing.T) {
	s := New(CharSets[10], 1*time.Second)
	delay1 := s.Delay
	s.UpdateSpeed(baseWait * time.Second)
	delay2 := s.Delay
	if delay1 == delay2 {
		t.Error("update of speed failed")
	}
	s = nil
}

// TestUpdateCharSet verifies that character sets can be updated
func TestUpdateCharSet(t *testing.T) {
	s := New(CharSets[14], 1*time.Second)
	charSet1 := s.chars
	s.UpdateCharSet(CharSets[1])
	charSet2 := s.chars
	for i := range charSet1 {
		if charSet1[i] == charSet2[i] {
			t.Error("update of char set failed")
		}
	}
	s = nil
}

// TestGenerateNumberSequence verifies that a string slice of a spefic size is returned
func TestGenerateNumberSequence(t *testing.T) {
	elementCount := 100
	seq := GenerateNumberSequence(elementCount)
	if reflect.TypeOf(seq).String() != "[]string" {
		t.Error("received incorrect type in return from GenerateNumberSequence")
	}
	t.Log("In: ", elementCount)
	t.Log("Out: ", len(seq))
	if len(seq) != elementCount {
		t.Error("number of elements in slice doesn't match expected count")
	}
}

// TestBackspace proves that the correct number of characters are removed.
func TestBackspace(t *testing.T) {
	// Because of buffering of output and time weirdness, somethings
	// are broken for an indeterminant reason without a wait
	time.Sleep(75 * time.Millisecond)
	fmt.Println()
	s := New(CharSets[0], 100*time.Millisecond)
	s.Color("blue")
	s.Start()
	fmt.Print("This is on the same line as the spinner: ")
	time.Sleep(baseWait * time.Second)
	s.Stop()
}

// TestColorError tests that if an invalid color string is passed to the Color
// function, the invalid color error is returned
func TestColorError(t *testing.T) {
	s := New(CharSets[0], 100*time.Millisecond)

	const invalidColorName = "bluez"
	const validColorName = "green"

	if s.Color(invalidColorName) != errInvalidColor {
		t.Error("Color method did not return an error when given an invalid color.")
	}

	if s.Color(validColorName) != nil {
		t.Error("Color method did not return nil when given a valid color name.")
	}
}

func TestWithWriter(t *testing.T) {
	s := New(CharSets[9], time.Millisecond*400, WithWriter(ioutil.Discard))
	_ = s
}

func TestWithNewLine(t *testing.T) {
	s := New(CharSets[9], 100*time.Millisecond, WithNewLine(true))
	if !s.NewLine {
		t.Error("expected NewLine to be true")
	}

	var buf syncBuffer
	s.Writer = &buf
	if f, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0); err == nil {
		defer f.Close()
		s.WriterFile = f
	}

	if !term.IsTerminal(int(s.WriterFile.Fd())) {
		t.Log("not running in a terminal, skipping Start execution")
		return
	}

	buf.WriteString("previous line output")
	s.Start()
	time.Sleep(150 * time.Millisecond)
	s.Stop()

	output := buf.String()
	expectedPrefix := "previous line output\n"
	if !strings.HasPrefix(output, expectedPrefix) {
		t.Errorf("expected output to start with %q, got %q", expectedPrefix, output)
	}
}

func TestWithNoClear(t *testing.T) {
	s := New(CharSets[9], 100*time.Millisecond, WithNoClear(true))
	if !s.NoClear {
		t.Error("expected NoClear to be true")
	}

	var buf syncBuffer
	s.Writer = &buf
	s.FinalMSG = "completed\n"

	if f, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0); err == nil {
		defer f.Close()
		s.WriterFile = f
	}

	if !term.IsTerminal(int(s.WriterFile.Fd())) {
		t.Log("not running in a terminal, skipping Start execution")
		return
	}

	buf.WriteString("task in progress: ")
	s.Start()
	time.Sleep(150 * time.Millisecond)
	s.Stop()

	output := buf.String()
	if !strings.HasPrefix(output, "task in progress: ") {
		t.Errorf("expected initial text to be preserved, got %q", output)
	}
	if strings.Contains(output, "\r") {
		t.Errorf("expected output to not contain carriage returns, got %q", output)
	}
	if !strings.HasSuffix(output, "completed\n") {
		t.Errorf("expected output to end with final message, got %q", output)
	}
}

func TestNoClearWithPrefixAndSuffix(t *testing.T) {
	s := New(CharSets[9], 50*time.Millisecond, WithNoClear(true))
	s.Prefix = "step: "
	s.Suffix = " running"

	var buf syncBuffer
	s.Writer = &buf

	if f, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0); err == nil {
		defer f.Close()
		s.WriterFile = f
	}

	if !term.IsTerminal(int(s.WriterFile.Fd())) {
		t.Log("not running in a terminal, skipping Start execution")
		return
	}

	s.Start()
	time.Sleep(120 * time.Millisecond)
	s.Stop()

	output := buf.String()
	if strings.Contains(output, "\r") {
		t.Errorf("expected output without carriage return, got %q", output)
	}
	if !strings.Contains(output, "step: ") {
		t.Errorf("expected output to contain prefix, got %q", output)
	}
	if !strings.Contains(output, " running") {
		t.Errorf("expected output to contain suffix, got %q", output)
	}
}

func TestEraseNoClear(t *testing.T) {
	s := New(CharSets[9], 100*time.Millisecond, WithNoClear(true))
	var buf bytes.Buffer
	s.Writer = &buf

	// When lastOutputPlain is empty, erase should write nothing
	s.Lock()
	s.lastOutputPlain = ""
	s.erase()
	s.Unlock()
	if buf.Len() != 0 {
		t.Errorf("expected no bytes written when lastOutputPlain is empty, got %d", buf.Len())
	}

	// When lastOutputPlain contains characters, erase should use backspaces
	buf.Reset()
	s.Lock()
	s.lastOutputPlain = "abc"
	s.erase()
	s.Unlock()

	expectedANSI := "\b\b\b\033[K"
	if buf.String() != expectedANSI {
		t.Errorf("expected erase to write %q, got %q", expectedANSI, buf.String())
	}
	if s.lastOutputPlain != "" {
		t.Errorf("expected lastOutputPlain to be cleared, got %q", s.lastOutputPlain)
	}
}

func TestEraseDefault(t *testing.T) {
	s := New(CharSets[9], 100*time.Millisecond)
	var buf bytes.Buffer
	s.Writer = &buf

	s.Lock()
	s.lastOutputPlain = "abc"
	s.erase()
	s.Unlock()

	expectedANSI := "\r\033[K"
	if buf.String() != expectedANSI {
		t.Errorf("expected erase to write %q, got %q", expectedANSI, buf.String())
	}
	if s.lastOutputPlain != "" {
		t.Errorf("expected lastOutputPlain to be cleared, got %q", s.lastOutputPlain)
	}
}

func TestComputeNumberOfLinesNeededToPrintStringInternal(t *testing.T) {
	tests := []struct {
		description   string
		expectedCount int
		printedLine   string
		maxLineWidth  int
	}{
		{"BlankLine", 1, "", 50},
		{"SingleLine", 1, "Hello world", 50},
		{"SingleLineANSI", 1, "Hello \x1b[36mworld\x1b[0m", 20},
		{"MultiLine", 2, "Hello\n world", 50},
		{"MultiLineANSI", 2, "Hello\n \x1b[1;36mworld\x1b[0m", 20},
		{"LongString", 2, "Hello world! I am a super long string that will be printed in 2 lines", 50},
		{"LongStringWithNewlines", 4, "Hello world!\nI am a super long string that will be printed in 2 lines.\nAnother new line", 50},
		{"NewlineCharAtStart", 2, "\nHello world!", 50},
		{"NewlineCharAtStartANSI", 2, "\n\x1b[36mHello\x1b[0m world!", 50},
		{"NewlineCharAtStartANSIFlipped", 2, "\x1b[36m\nHello\x1b[0m world!", 50},
		{"MultipleNewlineCharAtStart", 4, "\n\n\nHello world!", 50},
		{"NewlineCharAtEnd", 2, "Hello world!\n", 50},
		{"NewlineCharAtEndANSI", 2, "Hello \x1b[36mworld!\x1b[0m\n", 50},
		{"NewlineCharAtEndANSIFlipped", 2, "Hello \x1b[36mworld!\n\x1b[0m", 50},
		{"StringExactlySizeOfScreen", 1, strings.Repeat("a", 50), 50},
		{"StringExactlySizeOfScreenANSI", 1, "\x1b[36m" + strings.Repeat("a", 50), 50},
		{"StringOneGreaterThanSizeOfScreen", 2, strings.Repeat("a", 51), 50},
	}

	for _, test := range tests {
		result := computeNumberOfLinesNeededToPrintStringInternal(test.printedLine,
			test.maxLineWidth)
		if result != test.expectedCount {
			// Output error, resetting leftover ANSI sequences
			t.Errorf("%s: Line '%s\x1b[0m' shoud be printed on '%d' line, got '%d'",
				test.description, test.printedLine, test.expectedCount, result)
		}
	}
}

/*
Benchmarks
*/

// BenchmarkNew runs a benchmark for the New() function
func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(CharSets[1], 1*time.Second)
	}
}

func BenchmarkNewStartStop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := New(CharSets[1], 1*time.Second)
		s.Start()
		s.Stop()
	}
}
