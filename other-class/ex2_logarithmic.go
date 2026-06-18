package otherclass

// Ejercicio 2 — O(log n) Logarítmico
// Implementa BinarySearch que devuelve el índice de target o -1 si no existe.
// Complejidad objetivo: O(log n)

func BinarySearch(a []int, target int) int {
	low, high := 0, len(a)-1
	for low <= high {
		mid := (low + high) / 2
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
