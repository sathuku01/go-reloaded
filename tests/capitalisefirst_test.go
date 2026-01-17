package go_reloaded

import (
	"testing"
	ts "go_reloaded/pkg/capitalise"
	)

type CapitaliseTest struct {
	input  string
	output string
}



func TestCapitaliseFirstLetter(t *testing.T) {

	tests := []CapitaliseTest{
		{ input : "it (cap) was the best of times", output : "It was the best of times"},
		{input : "it (cap) was the best of times", output : "It was the best of times"},
	}

	for i, _ := range tests {
		got := ts.CapsFirstLetter(tests[i].input)
		want := tests[i].output
		if got != want {
			t.Errorf("got %s; want %s", got, want)
		}
	}
}