package problem824

import "testing"

type testUnit struct {
	input  string
	output string
}

func TestToGoatLatin(t *testing.T) {
	table := []testUnit{
		{
			"I speak Goat Latin",
			"Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			"The quick brown fox jumped over the lazy dog",
			"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}

	for _, tt := range table {
		if actual := toGoatLatin(tt.input); actual != tt.output {
			t.Errorf("toGoatLatin(%s), expect %s, but got %s", tt.input, tt.output, actual)
		}
	}
}
