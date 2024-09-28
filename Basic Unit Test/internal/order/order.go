package order

import (
	"github.com/Rhymond/go-money"
	"fmt"
)

type Order struct {
	ID                string
	CurrencyAlphaCode string
	Items             []Item
}

type Item struct {
	ID        string
	Quantity  uint
	UnitPrice *money.Money
}

func (o Order) ComputeTotal() (*money.Money, error) {
	amount := money.New(0, o.CurrencyAlphaCode)
	for _, item := range o.Items {
		var err error
		amount, err = amount.Add(item.UnitPrice.Multiply(int64(item.Quantity)))
		if err!= nil {
			return nil, fmt.Errorf("not adding item elements, error: %w", err)
		}
	}

	return amount, nil
}
