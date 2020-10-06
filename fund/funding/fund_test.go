package funding

import (
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkWithdrawals(b *testing.B) {
	if b.N < WORKERS {
		return
	}

	server := NewFundServer(b.N)

	dollarsPerFounder := b.N / WORKERS

	// WaitGroup structs don't need to be initialized
	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		// Let the waitgroup know we're adding a foroutine
		wg.Add(1)

		// Spawn off a founder worker, as a closure
		go func() {
			// Mark this worker done when the function finishes
			defer wg.Done()

			for i := 0; i < dollarsPerFounder; i++ {
				server.Commands <- WithdrawCommand{Amount: 1}
			}
		}() // Remember to call the closure!
	}

	// Wait for all the workers to finish
	wg.Wait()

	balanceReponseChan := make(chan int)
	server.Commands <- BalanceCommand{Response: balanceReponseChan}
	balance := <-balanceReponseChan

	if balance != 0 {
		b.Error("Balance wasn't zero:", balance)
	}
}
