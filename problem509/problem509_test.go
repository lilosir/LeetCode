package problem509

import "testing"

type testUnit struct {
	input  int
	output int
}

func TestFib(t *testing.T) {
	table := []testUnit{
		{
			2,
			1,
		},
		{
			3,
			2,
		},
		{
			4,
			3,
		},
	}

	for _, tt := range table {
		if actual := fib(tt.input); actual != tt.output {
			t.Errorf("fib(%d) expect %d, but got %d", tt.input, tt.output, actual)
		}
	}
}
