package problem1309

import "testing"

type testUnit struct {
	input  string
	output string
}

func TestFreqAlphabets(t *testing.T) {
	table := []testUnit{
		{
			"11111",
			"aaaaa",
		},
		{
			"10#11#12",
			"jkab",
		},
		{
			"1326#",
			"acz",
		},
		{
			"25#",
			"y",
		},
		{
			"12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#",
			"abcdefghijklmnopqrstuvwxyz",
		},
		{
			"26#11#418#5",
			"zkdre",
		},
	}

	for _, tt := range table {
		if actual := freqAlphabets(tt.input); actual != tt.output {
			t.Errorf("freqAlphabets(%s) expect %s, but got %s", tt.input, tt.output, actual)
		}
	}
}
