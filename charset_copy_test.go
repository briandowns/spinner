package spinner

import (
	"reflect"
	"testing"
	"time"
)

func TestCharSetInputsAreCopied(t *testing.T) {
	for _, useUpdate := range []bool{false, true} {
		name := "New"
		if useUpdate {
			name = "UpdateCharSet"
		}
		t.Run(name, func(t *testing.T) {
			input := []string{"A", "B", "C"}
			s := New(input, time.Second)
			if useUpdate {
				s = New([]string{"other"}, time.Second)
				s.UpdateCharSet(input)
			}
			input[0] = "changed"
			if !reflect.DeepEqual(s.chars, []string{"A", "B", "C"}) {
				t.Fatalf("caller mutation changed spinner: %q", s.chars)
			}
			s.Reverse()
			if !reflect.DeepEqual(input, []string{"changed", "B", "C"}) {
				t.Fatalf("Reverse changed caller input: %q", input)
			}
			if !reflect.DeepEqual(s.chars, []string{"C", "B", "A"}) {
				t.Fatalf("Reverse did not update spinner: %q", s.chars)
			}
		})
	}
}

func TestReverseDoesNotChangeOtherSpinner(t *testing.T) {
	input := []string{"A", "B", "C"}
	first := New(input, time.Second)
	second := New(input, time.Second)
	first.Reverse()
	if !reflect.DeepEqual(second.chars, []string{"A", "B", "C"}) {
		t.Fatalf("Reverse changed another spinner: %q", second.chars)
	}
}

func TestReverseDoesNotChangePreset(t *testing.T) {
	preset := append([]string(nil), CharSets[9]...)
	defer copy(CharSets[9], preset)
	s := New(CharSets[9], time.Second)
	s.Reverse()
	if !reflect.DeepEqual(CharSets[9], preset) {
		t.Fatalf("Reverse changed shared preset: %q", CharSets[9])
	}
}
