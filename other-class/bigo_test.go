package otherclass

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

// ============================================================================
// UNIT TESTS (To verify the implementations are correct)
// ============================================================================


func TestLookup(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}
	m := map[int]bool{10: true, 20: true, 30: true, 40: true, 50: true}

	t.Run("LookupInSlice", func(t *testing.T) {
		if !LookupInSlice(nums, 30) {
			t.Error("expected to find 30")
		}
		if LookupInSlice(nums, 99) {
			t.Error("did not expect to find 99")
		}
	})

	t.Run("LookupInMap", func(t *testing.T) {
		if !LookupInMap(m, 30) {
			t.Error("expected to find 30")
		}
		if LookupInMap(m, 99) {
			t.Error("did not expect to find 99")
		}
	})
}

func TestBinarySearch(t *testing.T) {
	sortedNums := []int{10, 20, 30, 40, 50, 60, 70}

	tests := []struct {
		target   int
		expected int
	}{
		{10, 0},
		{40, 3},
		{70, 6},
		{25, -1},
		{5, -1},
		{80, -1},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Target_%d", tc.target), func(t *testing.T) {
			idxLinear := LinearSearchSorted(sortedNums, tc.target)
			idxBinary := BinarySearchSorted(sortedNums, tc.target)
			if idxLinear != tc.expected {
				t.Errorf("LinearSearchSorted: expected index %d, got %d", tc.expected, idxLinear)
			}
			if idxBinary != tc.expected {
				t.Errorf("BinarySearchSorted: expected index %d, got %d", tc.expected, idxBinary)
			}
		})
	}
}

func TestContainsDuplicates(t *testing.T) {
	dup := []int{1, 2, 3, 4, 2, 5}
	unique := []int{1, 2, 3, 4, 5, 6}

	t.Run("BruteForce", func(t *testing.T) {
		if !ContainsDuplicatesBruteForce(dup) {
			t.Error("expected duplicates to be true")
		}
		if ContainsDuplicatesBruteForce(unique) {
			t.Error("expected duplicates to be false")
		}
	})

	t.Run("Efficient", func(t *testing.T) {
		if !ContainsDuplicatesEfficient(dup) {
			t.Error("expected duplicates to be true")
		}
		if ContainsDuplicatesEfficient(unique) {
			t.Error("expected duplicates to be false")
		}
	})
}

func TestHasTwoSum(t *testing.T) {
	nums := []int{10, 15, 3, 7}
	target := 17 // 10 + 7

	t.Run("BruteForce", func(t *testing.T) {
		if !HasTwoSumBruteForce(nums, target) {
			t.Error("expected target 17 to be true")
		}
		if HasTwoSumBruteForce(nums, 99) {
			t.Error("expected target 99 to be false")
		}
	})

	t.Run("Sort", func(t *testing.T) {
		// Create a copy of the input so the function under test doesn't mutate our test slice
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		if !HasTwoSumSort(testNums, target) {
			t.Error("expected target 17 to be true")
		}
		copy(testNums, nums)
		if HasTwoSumSort(testNums, 99) {
			t.Error("expected target 99 to be false")
		}
	})
}

func TestComputeDistanceMatrix(t *testing.T) {
	points := []Point{
		{X: 0, Y: 0},
		{X: 3, Y: 0},
		{X: 0, Y: 4},
	}

	matrix := ComputeDistanceMatrix(points)
	if matrix == nil {
		t.Fatal("expected non-nil matrix, got nil (implement the function first)")
	}
	if len(matrix) != 3 || len(matrix[0]) != 3 {
		t.Fatalf("expected 3x3 matrix, got %dx%d", len(matrix), len(matrix[0]))
	}

	// Distance from point 0 to point 1 (0,0 to 3,0) = 3
	if math.Abs(matrix[0][1]-3.0) > 1e-9 {
		t.Errorf("expected distance between 0 and 1 to be 3.0, got %f", matrix[0][1])
	}
	// Distance from point 1 to point 2 (3,0 to 0,4) = sqrt(3^2 + 4^2) = 5
	if math.Abs(matrix[1][2]-5.0) > 1e-9 {
		t.Errorf("expected distance between 1 and 2 to be 5.0, got %f", matrix[1][2])
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Fib_%d", tc.n), func(t *testing.T) {
			valRecursive := FibonacciRecursive(tc.n)
			valIterative := FibonacciIterative(tc.n)
			if valRecursive != tc.expected {
				t.Errorf("FibonacciRecursive(%d): expected %d, got %d", tc.n, tc.expected, valRecursive)
			}
			if valIterative != tc.expected {
				t.Errorf("FibonacciIterative(%d): expected %d, got %d", tc.n, tc.expected, valIterative)
			}
		})
	}
}

// ============================================================================
// BENCHMARKS (To measure scaling and performance)
// ============================================================================

// Helper to generate a slice of sequential integers
func makeSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i
	}
	return s
}

// Helper to generate a map of size N
func makeMap(size int) map[int]bool {
	m := make(map[int]bool, size)
	for i := 0; i < size; i++ {
		m[i] = true
	}
	return m
}

// Helper to generate Points
func makePoints(size int) []Point {
	p := make([]Point, size)
	for i := 0; i < size; i++ {
		p[i] = Point{X: float64(i), Y: float64(i)}
	}
	return p
}

// --- Benchmark Exercise 1: O(1) vs O(N) Lookup ---

func BenchmarkLookupSlice_10(b *testing.B) {
	nums := makeSlice(10)
	target := -1 // Worst case (not found)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LookupInSlice(nums, target)
	}
}

func BenchmarkLookupSlice_1000(b *testing.B) {
	nums := makeSlice(1000)
	target := -1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LookupInSlice(nums, target)
	}
}

func BenchmarkLookupSlice_100000(b *testing.B) {
	nums := makeSlice(100000)
	target := -1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LookupInSlice(nums, target)
	}
}

func BenchmarkLookupMap_10(b *testing.B) {
	m := makeMap(10)
	target := -1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LookupInMap(m, target)
	}
}

func BenchmarkLookupMap_1000(b *testing.B) {
	m := makeMap(1000)
	target := -1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LookupInMap(m, target)
	}
}

func BenchmarkLookupMap_100000(b *testing.B) {
	m := makeMap(100000)
	target := -1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LookupInMap(m, target)
	}
}

// --- Benchmark Exercise 2: O(log N) vs O(N) Search ---

func BenchmarkLinearSearch_10(b *testing.B) {
	nums := makeSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearchSorted(nums, -1)
	}
}

func BenchmarkLinearSearch_1000(b *testing.B) {
	nums := makeSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearchSorted(nums, -1)
	}
}

func BenchmarkLinearSearch_100000(b *testing.B) {
	nums := makeSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearchSorted(nums, -1)
	}
}

func BenchmarkBinarySearch_10(b *testing.B) {
	nums := makeSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearchSorted(nums, -1)
	}
}

func BenchmarkBinarySearch_1000(b *testing.B) {
	nums := makeSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearchSorted(nums, -1)
	}
}

func BenchmarkBinarySearch_100000(b *testing.B) {
	nums := makeSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearchSorted(nums, -1)
	}
}

// --- Benchmark Exercise 3: O(N) vs O(N^2) Duplicates ---

func BenchmarkContainsDuplicatesEfficient_10(b *testing.B) {
	nums := makeSlice(10) // Unique numbers (worst case)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicatesEfficient(nums)
	}
}

func BenchmarkContainsDuplicatesEfficient_100(b *testing.B) {
	nums := makeSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicatesEfficient(nums)
	}
}

func BenchmarkContainsDuplicatesEfficient_1000(b *testing.B) {
	nums := makeSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicatesEfficient(nums)
	}
}

func BenchmarkContainsDuplicatesBruteForce_10(b *testing.B) {
	nums := makeSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicatesBruteForce(nums)
	}
}

func BenchmarkContainsDuplicatesBruteForce_100(b *testing.B) {
	nums := makeSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicatesBruteForce(nums)
	}
}

func BenchmarkContainsDuplicatesBruteForce_1000(b *testing.B) {
	nums := makeSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicatesBruteForce(nums)
	}
}

// --- Benchmark Exercise 4: O(N log N) vs O(N^2) Two Sum ---

func BenchmarkHasTwoSumSort_10(b *testing.B) {
	nums := makeSlice(10)
	slices.Reverse(nums) // Start with unsorted data
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Use a local copy to avoid modifying data across iterations
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		HasTwoSumSort(testNums, -1)
	}
}

func BenchmarkHasTwoSumSort_100(b *testing.B) {
	nums := makeSlice(100)
	slices.Reverse(nums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		HasTwoSumSort(testNums, -1)
	}
}

func BenchmarkHasTwoSumSort_1000(b *testing.B) {
	nums := makeSlice(1000)
	slices.Reverse(nums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		HasTwoSumSort(testNums, -1)
	}
}

func BenchmarkHasTwoSumBruteForce_10(b *testing.B) {
	nums := makeSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasTwoSumBruteForce(nums, -1)
	}
}

func BenchmarkHasTwoSumBruteForce_100(b *testing.B) {
	nums := makeSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasTwoSumBruteForce(nums, -1)
	}
}

func BenchmarkHasTwoSumBruteForce_1000(b *testing.B) {
	nums := makeSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasTwoSumBruteForce(nums, -1)
	}
}

// --- Benchmark Exercise 5: O(N^2) Distance Matrix ---

func BenchmarkComputeDistanceMatrix_10(b *testing.B) {
	points := makePoints(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ComputeDistanceMatrix(points)
	}
}

func BenchmarkComputeDistanceMatrix_100(b *testing.B) {
	points := makePoints(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ComputeDistanceMatrix(points)
	}
}

func BenchmarkComputeDistanceMatrix_500(b *testing.B) {
	points := makePoints(500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ComputeDistanceMatrix(points)
	}
}

// --- Benchmark Exercise 6: O(2^N) vs O(N) Fibonacci ---

func BenchmarkFibonacciRecursive_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursive(10)
	}
}

func BenchmarkFibonacciRecursive_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursive(20)
	}
}

func BenchmarkFibonacciRecursive_30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursive(30)
	}
}

func BenchmarkFibonacciIterative_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIterative(10)
	}
}

func BenchmarkFibonacciIterative_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIterative(20)
	}
}

func BenchmarkFibonacciIterative_30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIterative(30)
	}
}

func BenchmarkFibonacciIterative_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIterative(10000)
	}
}
