# Desafío de Entrevista: Inversión de Punteros en Memoria (Linked List)

## Contexto
En entrevistas técnicas para roles intermedios y senior en Go, los evaluadores usan listas enlazadas no para ver si sabés armar una lista, sino para medir si dominás el **rastreo de punteros en memoria** sin generar ciclos infinitos ni perder referencias en el Heap.

## El Enunciado
Dada la cabeza (`head`) de una lista enlazada simple, invertí la lista **en memoria (in-place)** y retorná la nueva cabeza.

Restricciones:
- Complejidad temporal: **O(n)**.
- Complejidad espacial: **O(1)** (no podés crear nodos nuevos ni usar slices auxiliares; tenés que reorientar los punteros existentes).
