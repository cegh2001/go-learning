package otherclass

// Exercise 3: Linear Time O(N) vs. Quadratic Time O(N^2)
//
// Problem: Detect if a slice contains any duplicate numbers.
//
// This exercise highlights the classic trade-off between TIME complexity and SPACE complexity:
// - Brute Force uses O(N^2) time but O(1) extra space.
// - Efficient uses O(N) time but O(N) extra space by utilizing a hash map.

// ContainsDuplicatesBruteForce checks for duplicates by comparing every element with every other element.
// Time Complexity: O(N^2)
// Space Complexity: O(1)
func ContainsDuplicatesBruteForce(nums []int) bool {
	// TODO: Compare each element nums[i] with every other element nums[j] where i != j.
	return false
}

// ContainsDuplicatesEfficient checks for duplicates by keeping track of visited elements in a map.
// Time Complexity: O(N)
// Space Complexity: O(N)
func ContainsDuplicatesEfficient(nums []int) bool {
	// TODO: Use a map to remember numbers we have already seen.
	// If you see a number twice, return true.
	return false
}
