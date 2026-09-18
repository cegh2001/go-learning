package entrevista

import (
	"reflect"
	"testing"
)

func toSlice(head *ListNode) []int {
	var result []int
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func fromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "lista de 5 elementos",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "lista de 2 elementos",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "lista vacía",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := fromSlice(tc.input)
			gotHead := ReverseList(head)
			got := toSlice(gotHead)

			if len(tc.expected) == 0 && len(got) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("ReverseList(%v) = %v; se esperaba %v", tc.input, got, tc.expected)
			}
		})
	}
}
