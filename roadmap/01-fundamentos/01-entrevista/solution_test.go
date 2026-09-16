package entrevista

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "caso básico positivo",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "elementos dispersos",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "valores idénticos",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "incluye números negativos",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "cero en el conjunto",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)
			if len(got) != 2 {
				t.Fatalf("se esperaban 2 índices, se obtuvo: %v", got)
			}

			// Normalizar orden para comparar
			sort.Ints(got)
			sort.Ints(tc.expected)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("TwoSum(%v, %d) = %v; se esperaba %v", tc.nums, tc.target, got, tc.expected)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i * 2
	}
	// Target es la suma de los dos últimos
	target := nums[9998] + nums[9999]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}
