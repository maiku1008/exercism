package account

import "sync"

// Account represents a bank account
type Account struct {
	active  bool
	balance int64
	mu      sync.Mutex
}

// Open creates a new bank account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		active:  true,
		balance: initialDeposit,
	}
}

// Balance returns the balance of an account
func (a *Account) Balance() (balance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.active {
		return 0, false
	}
	return a.balance, true
}

// Deposit adds an amount of an account's balance
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.active {
		return 0, false
	}
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

// Close closes a bank account
func (a *Account) Close() (payout int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.active {
		return 0, false
	}
	a.active = false
	return a.balance, true
}
