package pointerserrors

type Wallet struct {
	bal int
}

func (w *Wallet) Deposit(amount int) {
	w.bal += amount
}

func (w Wallet) Balance() int {
	return w.bal
}
