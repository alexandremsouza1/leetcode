package src

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		s        []int
		t        int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 7, []int{3, 4}},
		{[]int{1, 2, 3, 4}, 5, []int{2, 3}},
		{[]int{0, 5, 15, 7}, 12, []int{5, 7}},
		{[]int{3, 6, 9, 12}, 9, []int{3, 6}},
	}

	for _, test := range tests {
		result := IsTwoSum(test.s, test.t)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("IsTwoSum(%v, %v) = %v; expected %v", test.s, test.t, result, test.expected)
		}
	}
}
