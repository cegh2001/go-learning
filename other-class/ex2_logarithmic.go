package otherclass

// Exercise 2: Logarithmic Time O(log N) vs. Linear Time O(N)
//
// Searching in a SORTED list. If we don't exploit the sorting, we search line-by-line (O(N)).
// If we divide and conquer, we get O(log N) via Binary Search.
//
// Goal: Implement both algorithms. Observe how O(log N) performs virtually the same number
// of steps whether N is 100 or 1,000,000!

// LinearSearchSorted searches the target in a sorted slice by inspecting every element sequentially.
// Complexity: O(N)
func LinearSearchSorted(sortedNums []int, target int) int {
	// TODO: Perform a linear search and return the index of target, or -1 if not found.
	return -1
}

// BinarySearchSorted searches the target in a sorted slice using the divide-and-conquer strategy.
// Complexity: O(log N)
func BinarySearchSorted(sortedNums []int, target int) int {
	// TODO: Implement binary search.
	// Keep track of low and high pointers, calculate the midpoint,
	// and narrow down the search space by half in each iteration.
	return -1
}
