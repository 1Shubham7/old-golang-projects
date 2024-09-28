package order

import (
	"testing"

	"github.com/Rhymond/go-money"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	o := Order{
		ID: "100",
		CurrencyAlphaCode: "INR",
		Items: []Item{
			{
			ID: "500",
			Quantity: 2,
			UnitPrice: money.New(100, "INR"),
		},
	},
}
	total, err := o.ComputeTotal()
	assert.NoError(t, err)
	assert.Equal(t, int64(200), total.Amount())
	assert.Equal(t, "INR", total.Currency().Code)
}