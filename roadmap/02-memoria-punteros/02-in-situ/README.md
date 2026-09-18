# Ticket JIRA: FINTECH-302 — Bug Crítico: Los saldos y auditorías no se actualizan

**Prioridad**: Bloqueante  
**Componente**: `core-banking / wallet-service`

---

## Descripción del Problema
El equipo de operaciones reporta que los usuarios están realizando depósitos y transferencias, pero **el saldo de la billetera nunca se incrementa** y el historial de transacciones queda vacío.

El desarrollador que lo implementó insiste: *"En mi máquina el código no tira errores de compilación ni panics, la suma matemática está perfecta"*.

## Tu Misión
1. Abrir `wallet.go` y revisar los métodos `Deposit`, `Withdraw` y `AddAuditLog`.
2. Explicar por qué Go descarta los cambios al ejecutarse (el concepto de **Value Receiver vs Pointer Receiver** y **copia de Stack**).
3. Refactorizar el código para que las mutaciones persistan en la billetera original.
4. Correr los tests para verificar que el bug esté erradicado:
   ```bash
   go test -v .
   ```
