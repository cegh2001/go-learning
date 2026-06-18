package otherclass

// Ejercicio 3 — O(n) Lineal
// Implementa Sum y Contains: sumar todos los elementos y comprobar existencia.
// Complejidad objetivo: O(n)

func Sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func Contains(nums []int, target int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}
	return false
}
