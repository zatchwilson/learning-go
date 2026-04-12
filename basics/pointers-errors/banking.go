package pointerserrors

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) { //Passing pointer
	w.balance += amount
}

func (w *Wallet) Balance() int { //Passing pointer to the wallet we need to alter
	return w.balance
}
