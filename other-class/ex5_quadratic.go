package otherclass

// Ejercicio 5 — O(n^2) Quadratic
// Dado un slice de enteros, devuelve si existe algún par (i,j) con la misma suma objetivo.
// Implementación naive con doble bucle.
// Complejidad objetivo: O(n^2)

func HasPairWithSum(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}
