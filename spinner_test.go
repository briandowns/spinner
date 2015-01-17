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
	"reflect"
	"testing"
	"time"
)

func withOutput(a []string, d time.Duration) (*Spinner, *bytes.Buffer) {
	s := New(a, d)
	out := new(bytes.Buffer)
	s.w = out
	return s, out
}

func byteSlicesEq(a []byte, b []byte) bool {
	return string(a) == string(b)
}

// TestNew verifies that the returned instance is of the proper type
func TestNew(t *testing.T) {
	s := New(CharSets[1], 1*time.Second)
	if reflect.TypeOf(s).String() != "*spinner.Spinner" {
		t.Error("New returned incorrect type")
	}
}

// TestStart will verify a spinner can be started
func TestStart(t *testing.T) {
	s := New(CharSets[1], 100*time.Millisecond)
	s.Color("red")
	s.Start()
	dur := 4000
	time.Sleep(time.Duration(dur) * time.Millisecond)
	s.Stop()
	time.Sleep(100 * time.Millisecond)
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
	len1 := out.Len()
	time.Sleep(300 * time.Millisecond)
	len2 := out.Len()
	if len1 != len2 {
		t.Errorf("expected equal, got %v != %v", len1, len2)
	}
	p = nil
}

// TestRestart will verify a spinner can be stopped and started again
func TestRestart(t *testing.T) {
	s := New(CharSets[4], 50*time.Millisecond)
	out := new(bytes.Buffer)
	s.w = out
	s.Start()
	s.Color("cyan")
	time.Sleep(200 * time.Millisecond)
	s.Restart()
	time.Sleep(200 * time.Millisecond)
	s.Stop()
	time.Sleep(50 * time.Millisecond)
	result := out.Bytes()
	first := result[:len(result)/2]
	secnd := result[len(result)/2:]
	if !byteSlicesEq(first, secnd) {
		t.Errorf("Expected ==, got \n%#v != \n%#v", first, secnd)
	}
	s = nil
}

// TestReverse will verify that the given spinner can stop and start again reversed
func TestReverse(t *testing.T) {
	a := New(CharSets[10], 1*time.Second)
	a.Color("red")
	a.Start()
	time.Sleep(4 * time.Second)
	a.Reverse()
	a.Restart()
	time.Sleep(4 * time.Second)
	a.Reverse()
	a.Restart()
	time.Sleep(4 * time.Second)
	a.Stop()
	a = nil
}

// TestUpdateSpeed verifies that the delay can be updated
func TestUpdateSpeed(t *testing.T) {
	s := New(CharSets[10], 1*time.Second)
	delay1 := s.Delay
	s.UpdateSpeed(3 * time.Second)
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
	if len(seq) != elementCount {
		t.Error("number of elements in slice doesn't match expected count")
	}
}

// TestMultiple will
func TestMultiple(t *testing.T) {
	a := New(CharSets[0], 100*time.Millisecond)
	b := New(CharSets[1], 250*time.Millisecond)
	a.Start()
	a.Color("green")
	b.Start()
	time.Sleep(4 * time.Second)
	a.Stop()
	time.Sleep(3 * time.Second)
	b.Stop()
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
	time.Sleep(4 * time.Second)
	s.Stop()
}
