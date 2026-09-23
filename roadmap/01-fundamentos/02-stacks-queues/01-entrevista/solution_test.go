package entrevista

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "paréntesis simples",
			input:    "()",
			expected: true,
		},
		{
			name:     "múltiples tipos correctos",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "anidados correctos",
			input:    "{[()]}",
			expected: true,
		},
		{
			name:     "cierre cruzado incorrecto",
			input:    "(]",
			expected: false,
		},
		{
			name:     "orden incorrecto",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "solo apertura",
			input:    "[",
			expected: false,
		},
		{
			name:     "solo cierre",
			input:    "]",
			expected: false,
		},
		{
			name:     "cadena vacía",
			input:    "",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValid(tc.input)
			if got != tc.expected {
				t.Errorf("IsValid(%q) = %v; se esperaba %v", tc.input, got, tc.expected)
			}
		})
	}
}
