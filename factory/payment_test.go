package factory

import (
	"strings"
	"testing"
)

func TestCashPayment_Pay(t *testing.T) {
	cashPayment, err := GetPaymentMethod(1)
	if err != nil {
		t.Errorf("Error finding payment factory method")
	}
	res := cashPayment.Pay(10.)
	if !strings.Contains(res, "cash") {
		t.Errorf("The cash payment method was not found")
	}
}
