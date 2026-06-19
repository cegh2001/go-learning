package otherclass

// Exercise 4: Linearithmic Time O(N log N) vs. Quadratic Time O(N^2)
//
// Problem: Check if there are two numbers in a slice that sum up to a target value.
//
// In this lab, we compare:
// 1. Brute force nested loops: O(N^2)
// 2. Sorting the slice first (O(N log N)) and then using a two-pointer approach (O(N)): Total: O(N log N)

// HasTwoSumBruteForce checks every possible pair in the slice to see if they sum up to target.
// Time Complexity: O(N^2)
// Space Complexity: O(1)
func HasTwoSumBruteForce(nums []int, target int) bool {
	// TODO: Use two nested loops to check if nums[i] + nums[j] == target.
	return false
}

// HasTwoSumSort sorts the slice first and then uses two pointers (left and right) to find the pair.
// Note: You can use Go's standard library `slices.Sort` or `sort.Ints` to sort the array.
// Time Complexity: O(N log N) (due to sorting) + O(N) (two-pointer search) = O(N log N)
// Space Complexity: O(1) if sorting in place, or O(N) if copying the array first.
func HasTwoSumSort(nums []int, target int) bool {
	// TODO: Create a copy of nums (to avoid modifying the original input in tests),
	// sort the copy using the standard library, and use a left pointer (at 0) and
	// right pointer (at len-1). Move them inward based on their sum.
	return false
}
