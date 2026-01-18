package yourpackage

import (
	ts "go_reloaded/process"
	"testing"
)

type testCase struct {
	input  string
	output string
}

func StringToSlice(s string) []string {
	result := make([]string, 0, len(s))
	for _, r := range s {
		result = append(result, string(r))
	}
	return result
}

func TestProcessData(t *testing.T) {
	tests := []testCase{
		{
			input:  "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			output: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			input:  "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			output: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			input:  "Don not be sad ,because sad backwards is das . And das not good",
			output: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			input:  "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			output: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
	}

	for i, tt := range tests {

		got := ts.ProcessData(StringToSlice(tt.input))
		want := StringToSlice(tt.output)

		if !EqualStringSlices(got, want) {
			t.Errorf(
				"test %d failed\ninput:  %q\ngot:    %q\nwant:   %q",
				i, tt.input, got, tt.output,
			)
		}
	}
}

func EqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
