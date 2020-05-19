package problem884

import (
	"testing"

	"github.com/lilosir/LeetCode/kit"
)

type input struct {
	a string
	b string
}

type testUnit struct {
	input  input
	output []string
}

func TestToGoatLatin(t *testing.T) {
	table := []testUnit{
		{
			input{
				"this apple is sweet",
				"this apple is sour",
			},
			[]string{"sweet", "sour"},
		},
		{
			input{
				"apple apple",
				"banana",
			},
			[]string{"banana"},
		},
	}

	for _, tt := range table {
		actual := uncommonFromSentences(tt.input.a, tt.input.b)
		iactual := make([]interface{}, len(actual))
		ioutput := make([]interface{}, len(tt.output))
		for i := range actual {
			iactual[i] = actual[i]
		}
		for i := range tt.output {
			ioutput[i] = tt.output[i]
		}
		if !kit.CheckTwoSliceHaveSameElements(iactual, ioutput) {
			t.Errorf("toGoatLatin(%s, %s), expect %v, but got %v", tt.input.a, tt.input.b, ioutput, iactual)
		}
	}
}
