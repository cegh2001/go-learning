# Ticket JIRA: PROD-1042 — Degradación severa de CPU en ingestión de eventos

**Prioridad**: Alta  
**Componente**: `event-ingestion / batch-processor`  
**Autor del reporte**: Equipo de SRE / Plataforma

---

## Descripción del Problema
En el último despliegue, un desarrollador utilizó un asistente de IA para implementar rápidamente la deduplicación de IDs de transacciones en lotes de eventos (`DeduplicateEventIDs`).

En tests locales con 10 o 20 IDs el código funcionaba y los tests pasaban. Sin embargo, al recibir un lote de producción de 50.000 IDs, el pod de Kubernetes entra en **High CPU Alert (100%)**, los tiempos de respuesta superan los 5 segundos y el servicio es reiniciado por timeout del health check.

## Código Actual
Revisá `service.go`. Notarás cómo se está deduplicando actualmente la lista de strings.

## Tu Tarea
1. **Auditar el código existente**: Explicar técnicamente por qué colapsa en términos de **complejidad temporal Big O** y qué pasa con la **memoria/slices** en cada iteración.
2. **Refactorizar `DeduplicateEventIDs`**: Mantener el comportamiento esperado (conservar el orden de primera aparición o retornar los IDs únicos) reduciendo drásticamente la complejidad temporal a **O(n)** y minimizando las alocaciones de memoria.
3. **Validar con benchmarks**:
   ```bash
   go test -bench=. -benchmem ./...
   ```
   El tiempo por operación debe pasar de milisegundos/segundos a pocos microsegundos.
