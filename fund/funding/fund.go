package funding

// Fund struct
type Fund struct {
	// balance is unexported (private)
	balance int
}

// NewFund A regular function returning pointer to a fund
func NewFund(initialBalance int) Fund {
	return Fund{
		balance: initialBalance,
	}
}

// Balance method start with a *receiver*
func (f *Fund) Balance() int {
	return f.balance
}

// Withdraw method
func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
