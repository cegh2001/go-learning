package insitu

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeduplicateEventIDs(t *testing.T) {
	input := []string{"evt-1", "evt-2", "evt-1", "evt-3", "evt-2", "evt-4"}
	expected := []string{"evt-1", "evt-2", "evt-3", "evt-4"}

	got := DeduplicateEventIDs(input)
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("se esperaba %v, pero se obtuvo %v", expected, got)
	}
}

func BenchmarkDeduplicateEventIDs(b *testing.B) {
	// Simulamos un lote moderado de 10.000 IDs con repetidos
	totalEvents := 10000
	events := make([]string, totalEvents)
	for i := 0; i < totalEvents; i++ {
		// Genera repetidos cada 1000 eventos
		events[i] = fmt.Sprintf("txn-%d", i%1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DeduplicateEventIDs(events)
	}
}
