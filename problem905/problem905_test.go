package problem905

import "testing"

type testUnit struct {
	input []int
}

func TestSortArrayByParity(t *testing.T) {
	table := []testUnit{
		{
			[]int{3,1,2,4},
		},
		{
			[]int{3,1,2,4,7,4,6,3,11},
		},
	}

	for _, tt := range table {
		if actual := sortArrayByParity(tt.input); !isEvenFollowedByOdd(actual) {
			t.Errorf("sortArrayByParity(%v), got %v", tt.input, actual)
		}
	}
}

func isEvenFollowedByOdd(a []int) bool {
	index := 0
	for i, v := range a {
		if v%2 != 0 {
			index = i
			break;
		}
	}
	for i:=index; i<len(a); i++ {
		if a[i]%2 == 0 {
			return false
		}
	}
	return true
}