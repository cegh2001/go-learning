# Desafío de Entrevista: Two Sum & Anatomía de Memoria

## Contexto
Estás en una entrevista técnica para una posición de Backend Engineer en Go. Te piden resolver un problema clásico, pero no solo les interesa que el código compile: van a evaluar tu comprensión de **complejidad temporal, complejidad espacial y cómo Go maneja la memoria**.

## El Enunciado
Dado un slice de enteros `nums` y un entero `target`, devolvé los **índices** de los dos números que sumen `target`.

### Reglas y Restricciones
1. Podés asumir que cada entrada tiene **exactamente una solución**.
2. No podés usar el mismo elemento dos veces.
3. El orden de los índices en la respuesta no importa.
4. Complejidad temporal esperada: **O(n)**.
5. Complejidad espacial esperada: **O(n)** o mejor.

## Tu Misión
Implementá la función en `solution.go`:
```go
package entrevista

func TwoSum(nums []int, target int) []int {
    // Tu implementación acá
}
```

Luego corré los tests con:
```bash
go test -v ./...
```
