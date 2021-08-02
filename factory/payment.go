package factory

import (
	"errors"
	"fmt"
)

type PaymentFactory interface {
	Pay(amount float32) string
}

const(
	Cash = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentFactory, error){
	switch m {
	case Cash:
		return new(CashPayment), nil
	case DebitCard:
		return new(DebitCardPayment), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", m))
	}
}

type CashPayment struct {}
type DebitCardPayment struct {}

func (c CashPayment) Pay(amount float32) string  {
	return fmt.Sprintf("You payed using cash for %.2f", amount)
}

func (d DebitCardPayment) Pay(amount float32) string  {
	return fmt.Sprintf("You payed using debit card for %.2f", amount)
}
