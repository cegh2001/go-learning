package otherclass

// Ejercicio 6 — O(2^n) Exponencial
// Generar todos los subconjuntos (power set) de un slice de enteros.
// Complejidad objetivo: O(2^n) tiempo y O(2^n) espacio para almacenar resultados.

func PowerSet(nums []int) [][]int {
	res := [][]int{{}}
	for _, v := range nums {
		n := len(res)
		for i := 0; i < n; i++ {
			// crear nueva subset añadiendo v
			cur := make([]int, len(res[i])+1)
			copy(cur, res[i])
			cur[len(cur)-1] = v
			res = append(res, cur)
		}
	}
	return res
}
