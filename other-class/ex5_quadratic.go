package otherclass

// Exercise 5: Quadratic Time O(N^2)
//
// In this exercise, we practice O(N^2) complexity with a real scenario:
// computing a distance matrix between a list of 2D points.
//
// Because we must compute the distance between every possible pair of points,
// the time complexity is O(N^2) and the space complexity is also O(N^2) (to store the matrix).

// Point represents a coordinate in 2D space.
type Point struct {
	X, Y float64
}

// ComputeDistanceMatrix calculates the Euclidean distance between every pair of points.
// For N points, it returns an N x N matrix (slice of slices) where matrix[i][j] is the distance
// between points[i] and points[j].
// Time Complexity: O(N^2)
// Space Complexity: O(N^2)
func ComputeDistanceMatrix(points []Point) [][]float64 {
	// TODO: Initialize a 2D slice of size N x N.
	// Use nested loops to compute the Euclidean distance between points[i] and points[j].
	// Hint: Distance = math.Sqrt((x2-x1)^2 + (y2-y1)^2)
	// You can import the "math" package if needed.
	return nil
}
