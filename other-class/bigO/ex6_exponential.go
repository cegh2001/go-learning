package otherclass

// Exercise 6: Exponential Time O(2^N) vs. Linear Time O(N)
//
// The Fibonacci sequence is a classic example of how recursion can lead to exponential complexity.
//
// Goal: Implement the naive recursive Fibonacci (O(2^N)) and the iterative Fibonacci (O(N)).
//
// Warning: Running the benchmark for recursive Fibonacci with N >= 40 can lock up your terminal.
// That is why the benchmarks will use small values of N to keep it safe!

// FibonacciRecursive calculates the N-th Fibonacci number using naive recursion.
// Formula: F(n) = F(n-1) + F(n-2) with F(0) = 0, F(1) = 1.
// Time Complexity: O(2^N)
// Space Complexity: O(N) (due to call stack depth)
func FibonacciRecursive(n int) int {
	// TODO: Implement the naive recursive solution.
	// Base cases: if n <= 1, return n.
	// Recursive case: return FibonacciRecursive(n-1) + FibonacciRecursive(n-2).
	return 0
}

// FibonacciIterative calculates the N-th Fibonacci number using a simple loop.
// Time Complexity: O(N)
// Space Complexity: O(1)
func FibonacciIterative(n int) int {
	// TODO: Implement the iterative solution using a loop.
	// Keep track of the last two values and update them.
	return 0
}
