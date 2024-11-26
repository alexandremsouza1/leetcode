package src

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"listen", "silent", true},  // Caso clássico de anagrama
		{"triangle", "integral", true}, // Outro exemplo válido
		{"hello", "world", false},  // Não são anagramas
		{"test", "tset", true},     // Mesmo conjunto de caracteres
		{"anagram", "nagaram", true}, // Um caso mais longo
		{"rat", "car", false},      // Tamanhos diferentes
		{"aabbcc", "ccbbaa", true},   
		{"aabbcc", "aabbcd", true},   
	}

	for _, test := range tests {
		result := IsAnagram(test.s, test.t)
		if result != test.expected {
			t.Errorf("IsAnagram(%q, %q) = %v; expected %v", test.s, test.t, result, test.expected)
		}
	}
}
