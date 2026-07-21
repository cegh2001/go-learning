package otherclass

// Exercise 1: Constant Time O(1) vs. Linear Time O(N)
//
// In this exercise, we compare searching for an element in a list (slice)
// versus searching for an element in a hash map.
//
// Goal: Implement both functions and notice the huge performance difference in benchmarks.

// LookupInSlice searches for a target integer by traversing the slice from start to finish.
// Worst-case complexity: O(N) - Linear Time
func LookupInSlice(nums []int, target int) bool {
	// TODO: Iterate through the slice and return true if target is found, false otherwise.
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

// LookupInMap checks if a target integer exists as a key in the given map.
// Average-case complexity: O(1) - Constant Time
func LookupInMap(lookupMap map[int]bool, target int) bool {
	// TODO: Retrieve the boolean value directly from the map in O(1) time.
	if _, exists := lookupMap[target]; exists {
		return true
	}
	return false
}
