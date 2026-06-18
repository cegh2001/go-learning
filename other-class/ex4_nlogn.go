package otherclass

import "sort"

// Ejercicio 4 — O(n log n) Linearithmic
// Dado un slice de enteros, devuelve una copia ordenada y el número de valores únicos.
// Complejidad objetivo: O(n log n) debido al sort.

func SortedAndUniqueCount(nums []int) ([]int, int) {
	if len(nums) == 0 {
		return []int{}, 0
	}
	cop := make([]int, len(nums))
	copy(cop, nums)
	sort.Ints(cop)
	unique := 1
	for i := 1; i < len(cop); i++ {
		if cop[i] != cop[i-1] {
			unique++
		}
	}
	return cop, unique
}
