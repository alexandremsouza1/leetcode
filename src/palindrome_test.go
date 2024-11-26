package src

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"Was it a car or a cat I saw?",true},
		{"No lemon, no melon.",true},
		{"", true},
		{"Eva, can I see bear in a cave?",false},
		{"Mr. Owl ate chick my metal worm.",false},
		{"A Santa lived as a devil on NASA.",false},

	}

	for _, test := range tests {
		result := IsPalindrome(test.s)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; expected %v", test.s, result, test.expected)
		}
	}
}
