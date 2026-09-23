# Desafío de Entrevista: Validador de Paréntesis con Pila (Stack - LIFO)

## Contexto
En compiladores, parsers de JSON y procesadores de texto, las **Pilas (Stacks - LIFO: Last-In, First-Out)** son la estructura central.

## El Enunciado
Dada una cadena `s` que contiene únicamente los caracteres `(`, `)`, `{`, `}`, `[` y `]`, determiná si la cadena es válida.

Una cadena es válida si:
1. Los corchetes abiertos se cierran con el mismo tipo de corchetes.
2. Los corchetes abiertos se cierran en el orden correcto.
3. Cada corchete de cierre tiene un corchete abierto correspondiente del mismo tipo.

Restricciones:
- Complejidad temporal esperada: **O(n)**.
- Complejidad espacial esperada: **O(n)** (para la pila).
