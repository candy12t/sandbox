package money

import "fmt"

type Money struct {
	Amount   int
	Currency string
}

func New(amount int, currency string) (*Money, error) {
	if amount < 0 {
		return nil, fmt.Errorf("Argument Error: specify amount more than 0")
	}

	if len(currency) == 0 {
		return nil, fmt.Errorf("Argument Error: specify currency")
	}

	return &Money{
		Amount:   amount,
		Currency: currency,
	}, nil
}

func (m Money) Add(other Money) (*Money, error) {
	if m.Currency != other.Currency {
		return nil, fmt.Errorf("Argument Error: difference currency")
	}
	added := m.Amount + other.Amount
	return New(added, m.Currency)
}
