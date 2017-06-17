package example

import (
	"strconv"
	"testing"
)

////////////////////////////////////////////
// Paypal gateway mock implementation
////////////////////////////////////////////
type MockPaypal struct{}

// pay method implementation for the paypal gateway mock implementation
func (mp *MockPaypal) Pay(amount float32) string {
	value := strconv.FormatFloat(float64(amount), 'f', 2, 32)

	return "Called with value " + value
}

func TestShoppingCart_Checkout(t *testing.T) {
	t.Run("Testing shopping cart checkout feature", func(t *testing.T) {
		shoppingCart := ShoppingCart{
			payment: new(MockPaypal),
		}

		if shoppingCart.Checkout(100.00) == "USD: 100.00 Payment has been accepted by the paypal gateway" {
			t.Error("Pay method can't be called from Paypal struct")
		}

		if shoppingCart.Checkout(100.00) != "Called with value 100.00" {
			t.Error("Pay method must be called from MockPaypal struct")
		}
	})
}
