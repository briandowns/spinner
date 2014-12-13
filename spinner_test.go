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

package spinner

import (
	"reflect"
	"testing"
	"time"
)

// TestNew verifies that the returned instance is of the proper type
func TestNew(t *testing.T) {
	s := New(CharSets[10], 1*time.Second)
	if reflect.TypeOf(s).String() != "*spinner.Spinner" {
		t.Error("New returned incorrect type")
	}
}

// TestStart will verify a spinner can be started
func TestStart(t *testing.T) {
	s := New(CharSets[15], 1*time.Second)
	s.Start()
	time.Sleep(3 * time.Second)
	s.Stop()
}

// TestStop will verify a spinner can be stopped
func TestStop(t *testing.T) {
	s := New(CharSets[3], 1*time.Second)
	s.Start()
	time.Sleep(2 * time.Second)
	s.Stop()
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

// TestUpdateDelay verifies that the delay can be updated
func TestUpdateDelay(t *testing.T) {
	s := New(CharSets[8], 1*time.Second)
	delay1 := s.Delay
	s.UpdateDelay((3 * time.Second))
	delay2 := s.Delay
	if delay1 == delay2 {
		t.Error("update of delay set failed")
	}
}

// TestUpdateCharSet verifies that character sets can be updated
func TestUpdateCharSet(t *testing.T) {
	s := New(CharSets[15], 1*time.Second)
	charSet1 := s.Chars
	s.UpdateCharSet(CharSets[2])
	charSet2 := s.Chars
	for i, _ := range charSet1 {
		if charSet1[i] == charSet2[i] {
			t.Error("update of char set failed")
		}
	}
}
