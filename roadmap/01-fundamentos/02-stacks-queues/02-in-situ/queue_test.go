package insitu

import (
	"runtime"
	"testing"
)

func TestTaskQueue_FIFO(t *testing.T) {
	q := NewTaskQueue()
	t1 := &Task{ID: "1", Payload: []byte("task-1")}
	t2 := &Task{ID: "2", Payload: []byte("task-2")}

	q.Enqueue(t1)
	q.Enqueue(t2)

	got1, err := q.Dequeue()
	if err != nil || got1.ID != "1" {
		t.Fatalf("se esperaba tarea 1, se obtuvo: %v", got1)
	}

	got2, err := q.Dequeue()
	if err != nil || got2.ID != "2" {
		t.Fatalf("se esperaba tarea 2, se obtuvo: %v", got2)
	}

	if q.Len() != 0 {
		t.Fatalf("la cola debería estar vacía, pero tiene longitud %d", q.Len())
	}
}

func BenchmarkQueueLeak(b *testing.B) {
	q := NewTaskQueue()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Simula un flujo continuo de encolar y desencolar
		q.Enqueue(&Task{ID: "work", Payload: make([]byte, 1024)})
		_, _ = q.Dequeue()
	}

	// Forzamos al GC para ver el comportamiento
	runtime.GC()
}
