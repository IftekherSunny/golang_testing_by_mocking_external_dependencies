package golang_testing_by_mocking_external_dependencies

import "strconv"

////////////////////////////////////////////
// Payment method contract
////////////////////////////////////////////
type PaymentMethod interface {
	Pay(amount float32) string
}

////////////////////////////////////////////
// Paypal gateway implementation
////////////////////////////////////////////
type Paypal struct{}

// pay method implementation for the paypal gateway
func (s *Paypal) Pay(amount float32) string {

	// call gateway api and process billing
	usd := strconv.FormatFloat(float64(amount), 'f', 2, 32)

	return "USD: " + usd + " Payment has been accepted by the paypal gateway"
}

////////////////////////////////////////////
// Shopping cart
////////////////////////////////////////////
type ShoppingCart struct {
	payment PaymentMethod
}

// accepting shopping cart payment
func (s *ShoppingCart) Checkout(amount float32) string {
	return s.payment.Pay(amount)
}
