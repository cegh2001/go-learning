package insitu

import "errors"

type Task struct {
	ID      string
	Payload []byte
}

type TaskQueue struct {
	tasks []*Task
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks: make([]*Task, 0),
	}
}

// Enqueue agrega una tarea al final de la cola (FIFO).
func (q *TaskQueue) Enqueue(t *Task) {
	q.tasks = append(q.tasks, t)
}

// Dequeue extrae la tarea más antigua.
// PELIGRO: Reslicing ingenuo que causa retención de memoria.
func (q *TaskQueue) Dequeue() (*Task, error) {
	if len(q.tasks) == 0 {
		return nil, errors.New("cola vacía")
	}

	task := q.tasks[0]
	q.tasks[0] = nil // Evitar memory leak al liberar la referencia al primer elemento
	q.tasks = q.tasks[1:] // Causa memory leak de punteros en el array subyacente
	
	if len(q.tasks) == 0 {
		q.tasks = nil // Evitar memory leak al liberar todas las referencias al array subyacente
	}
	return task, nil
}

func (q *TaskQueue) Len() int {
	return len(q.tasks)
}
