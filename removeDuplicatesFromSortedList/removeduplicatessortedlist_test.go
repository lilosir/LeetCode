package problem83

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testUnit struct {
	input
	output
}

type input []int
type output []int

func TestRemoveDuplicates(t *testing.T) {
	assert := assert.New(t)

	units := []testUnit{
		testUnit{
			input:  []int{},
			output: []int{},
		},
		testUnit{
			input:  []int{1, 2, 4},
			output: []int{1, 2, 4},
		},
		testUnit{
			input:  []int{1, 2, 2, 3, 3, 5},
			output: []int{1, 2, 3, 5},
		},
	}

	for _, u := range units {
		input, output := u.input, u.output
		result := l2s(removeDuplicates(s2l(input)))
		// if result != output {
		// 	t.Errorf("Expected %v, but get %v", output, result)
		// }
		assert.EqualValues(output, result, "Expected %v, but get %v", output, result)
	}
}

func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
