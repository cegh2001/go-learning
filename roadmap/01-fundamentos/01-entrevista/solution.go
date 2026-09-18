package entrevista

// TwoSum encuentra dos índices en nums cuyos valores sumen target.
// Debe cumplir con complejidad temporal O(n) y espacial O(n).
func TwoSum(nums []int, target int) []int {
	// TODO: Implementar solución
	seen := make(map[int]int, len(nums)) // preasignar capacidad al map también mejora el rendimiento

	for index, number := range nums {
		complement := target - number
		if prev, exist := seen[complement]; exist {
			return []int{prev, index}
		}
		seen[number] = index
	}
	return nil
}
