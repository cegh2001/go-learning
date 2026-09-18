package insitu

import "errors"

type Wallet struct {
	Owner   string
	Balance float64
	Logs    []string
}

// NewWallet crea una nueva billetera
func NewWallet(owner string, initialBalance float64) Wallet {
	return Wallet{
		Owner:   owner,
		Balance: initialBalance,
		Logs:    make([]string, 0, 10),
	}
}

// Deposit suma saldo a la billetera.
// BUG: El saldo nunca cambia para quien llama a esta función.
func (w Wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("el monto debe ser mayor a cero")
	}

	w.Balance += amount
	w.Logs = append(w.Logs, "depósito realizado")
	return nil
}

// Withdraw descuenta saldo de la billetera.
// BUG: El saldo tampoco cambia acá.
func (w Wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("el monto debe ser mayor a cero")
	}
	if w.Balance < amount {
		return errors.New("fondos insuficientes")
	}

	w.Balance -= amount
	w.Logs = append(w.Logs, "retiro realizado")
	return nil
}
