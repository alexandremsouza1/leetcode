package src

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		s        []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 2}, // 2 e 3 são primos
		{[]int{5, 6, 7, 8}, 2}, // 5 e 7 são primos
		{[]int{10, 11, 12}, 1}, // Apenas 11 é primo

	}

	for _, test := range tests {
		result := IsCountPrimes(test.s)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; expected %v", test.s, result, test.expected)
		}
	}
}
