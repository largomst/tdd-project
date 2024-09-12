package main

import (
	"testing"
)

func TestMultiplicationDollar(t *testing.T) {
	five := Money{
		amount:   5,
		currency: "USD",
	}
	actualResult := five.Times(2)
	expectedResult := Money{
		amount:   10,
		currency: "USD",
	}
	if actualResult.amount != 10 {
		t.Errorf("Expected [%+v], got: [%+v]", expectedResult, actualResult)
	}
}
func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	actualResult := tenEuros.Times(2)
	expectedResult := Money{amount: 20, currency: "EUR"}
	if expectedResult.amount != actualResult.amount {
		t.Errorf("Expected [%+v], got: [%+v]", expectedResult, actualResult)
	}
}

func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}
	if expectedMoneyAfterDivision != actualMoneyAfterDivision {
		t.Errorf("Expected %+v Got %+v", expectedMoneyAfterDivision, actualMoneyAfterDivision)
	}
}

type Money struct {
	amount float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{m.amount * float64(multiplier), m.currency}
}


func (m Money) Divide(divisor int) Money {
	return Money{m.amount / float64(divisor), m.currency}
}