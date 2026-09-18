package insitu

import "testing"

func TestWallet_Deposit(t *testing.T) {
	w := NewWallet("Satoshi", 100.0)

	err := w.Deposit(50.0)
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}

	// El saldo debería ser 150.0
	if w.Balance != 150.0 {
		t.Errorf("FALLO: Balance esperado 150.0, obtenido: %.2f (¿se perdió la mutación en memoria?)", w.Balance)
	}

	// Debe haberse registrado el log
	if len(w.Logs) != 1 {
		t.Errorf("FALLO: Se esperaba 1 log de auditoría, obtenidos: %d", len(w.Logs))
	}
}

func TestWallet_Withdraw(t *testing.T) {
	w := NewWallet("Satoshi", 100.0)

	err := w.Withdraw(40.0)
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}

	if w.Balance != 60.0 {
		t.Errorf("FALLO: Balance esperado 60.0, obtenido: %.2f", w.Balance)
	}
}
