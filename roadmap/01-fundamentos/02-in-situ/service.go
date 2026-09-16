package insitu

// DeduplicateEventIDs toma una lista de IDs de eventos recibidos y retorna
// una lista sin duplicados preservando el orden original de primera aparición.
//
// ALERTA DE PRODUCCIÓN: Esta función colapsa en lotes de tamaño real.
// Debés auditarla y refactorizarla.
func DeduplicateEventIDs(events []string) []string {
	// var unique []string no es eficiente para lotes grandes, mejor preasignar capacidad
	unique := make([]string, 0, len(events))
	//seen := make(map[string]struct{})
	seen := make(map[string]struct{}, len(events)) // preasignar capacidad al map también mejora el rendimiento

	for _, id := range events {
		if _, exists := seen[id]; !exists {
			unique = append(unique, id)
			seen[id] = struct{}{} // es como hacer seen[id] = true
		}
	} //mejor big O en lugar de O(n^2) tenemos O(n) gracias al map
	
	// for _, id := range events {
	// 	// Busca linealmente si ya está en el slice resultante
	// 	exists := false
	// 	for _, u := range unique {
	// 		if u == id {
	// 			exists = true
	// 			break
	// 		}
	// 	}

	// 	if !exists {
	// 		unique = append(unique, id)
	// 	}
	// }

	return unique
}
