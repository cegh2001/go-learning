package insitu

// DeduplicateEventIDs toma una lista de IDs de eventos recibidos y retorna
// una lista sin duplicados preservando el orden original de primera aparición.
//
// ALERTA DE PRODUCCIÓN: Esta función colapsa en lotes de tamaño real.
// Debés auditarla y refactorizarla.
func DeduplicateEventIDs(events []string) []string {
	var unique []string

	for _, id := range events {
		// Busca linealmente si ya está en el slice resultante
		exists := false
		for _, u := range unique {
			if u == id {
				exists = true
				break
			}
		}

		if !exists {
			unique = append(unique, id)
		}
	}

	return unique
}
