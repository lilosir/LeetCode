/**
 * a little bit difficult to test, since the length of array/slice
 */

// package problem100

// import (
// 	"testing"
// )

// type treeVal interface{}

// type data struct {
// 	input  []treeVal
// 	output []treeVal
// 	result bool
// }

// func TestSameTree(t *testing.T) {
// 	tables := []data{
// 		data{
// 			[]treeVal{1, 3, 5},
// 			[]treeVal{1, 3, 5},
// 			true,
// 		},
// 		data{
// 			[]treeVal{1, 3},
// 			[]treeVal{1, 3, 5},
// 			false,
// 		},
// 		data{
// 			[]treeVal{1, 3, 5},
// 			[]treeVal{1, 5},
// 			false,
// 		},
// 		data{
// 			[]treeVal{1, nil, 5},
// 			[]treeVal{1, 3, 5},
// 			false,
// 		},
// 		data{
// 			nil,
// 			[]treeVal{1, 3, 5},
// 			false,
// 		},
// 		data{
// 			nil,
// 			nil,
// 			true,
// 		},
// 	}

// 	for _, t := range tables {
// 		input := t.input
// 	}
// }

// func s2t(v []treeVal) *TreeNode {

// }
