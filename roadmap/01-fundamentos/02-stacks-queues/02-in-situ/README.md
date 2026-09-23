# Ticket JIRA: INFRA-405 — Fuga sutil de memoria en cola de procesamiento en segundo plano

**Prioridad**: Alta  
**Componente**: `worker-pool / task-queue`

---

## El Problema
Un desarrollador implementó una cola FIFO (First-In, First-Out) para encolar y desencolar tareas en segundo plano.

La lógica funcional parece impecable:
- `Enqueue(task)`: agrega al final del slice.
- `Dequeue()`: saca el primer elemento haciendo `q.tasks = q.tasks[1:]`.

Sin embargo, en producción el servicio consume cada vez más megabytes de memoria RAM a lo largo de las horas y **nunca baja**, a pesar de que la cola casi siempre tiene 0 o 1 tareas pendientes.

## Tu Misión
1. Abrir `queue.go` y entender cómo funciona el operador de *reslicing* `tasks[1:]`.
2. Explicar por qué `tasks[1:]` no le permite al Garbage Collector liberar los objetos procesados.
3. Refactorizar la cola para que libere la referencia a los punteros procesados o implementar un mecanismo seguro.
4. Validar con tests:
   ```bash
   go test -v .
   ```
