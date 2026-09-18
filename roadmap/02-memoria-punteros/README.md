# Nivel 2: Gestión de Memoria, Punteros y Stack vs Heap

En este nivel nos metemos con las entrañas del runtime de Go:
- **Todo en Go se pasa por valor (Copia)**: No existe el paso por referencia nativo como en otros lenguajes; lo que pasás son copias (de valores o de punteros).
- **Punteros (`*T`, `&x`)**: Cuándo mutar estado y cuándo evitar fugas al Heap.
- **Stack vs Heap**: Entender qué variables son efímeras y cuáles fuerzan al Garbage Collector a trabajar.
