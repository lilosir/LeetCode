package problem1342

import (
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	testTable := []struct{ a, b int }{
		{14, 6},
		{8, 4},
		{0, 0},
		{1, 1},
		{123, 12},
	}

	for _, tt := range testTable {
		if actual := numberOfSteps(tt.a); actual != tt.b {
			t.Errorf("numberOfSteps(%d) expected %d, but got %d", tt.a, tt.b, actual)
		}
	}

}
