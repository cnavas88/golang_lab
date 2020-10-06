package funding

import "fmt"

// FundServer struct
type FundServer struct {
	Commands chan interface{}
	fund     Fund
}

// WithdrawCommand struct
type WithdrawCommand struct {
	Amount int
}

// BalanceCommand struct
type BalanceCommand struct {
	Response chan int
}

// NewFundServer create server
func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		Commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	// Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundServer) loop() {
	for command := range s.Commands {
		// command is an interface{}, but we can check its real type
		switch command.(type) {
		case WithdrawCommand:
			withdrawal := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)

		case BalanceCommand:
			getBalance := command.(BalanceCommand)
			balance := s.fund.Balance()
			getBalance.Response <- balance

		default:
			panic(fmt.Sprintf("Unrecognized command: %v", command))
		}
	}
}
