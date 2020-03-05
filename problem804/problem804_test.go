package problem804

import "testing"

type testUnit struct {
	input  []string
	output int
}

func TestUniqueMorseRepresentations(t *testing.T) {
	table := []testUnit{
		{
			[]string{"gin", "zen", "gig", "msg"},
			2,
		},
	}

	for _, tt := range table {
		if actual := uniqueMorseRepresentations(tt.input); actual != tt.output {
			t.Errorf("uniqueMorseRepresentations(%v) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
