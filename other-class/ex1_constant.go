package otherclass

// Ejercicio 1 — O(1) Constant
// Implementa GetFirst que devuelve el primer elemento de un slice.
// Complejidad objetivo: O(1)

func GetFirst(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	// TODO: devolver el primer elemento en tiempo constante
	return nums[0], true
}
