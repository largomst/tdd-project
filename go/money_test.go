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
	assertEqual(t, expectedResult, actualResult)
}
func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	actualResult := tenEuros.Times(2)
	expectedResult := Money{amount: 20, currency: "EUR"}
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioDollars Money

	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	fifteenDollars := Money{amount: 15, currency: "USD"}

	portfolio = portfolio.Add(fiveDollars, tenDollars)
	portfolioDollars = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portfolioDollars)
}

func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{m.amount * float64(multiplier), m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{m.amount / float64(divisor), m.currency}
}

type Portfolio []Money

func (p Portfolio) Add(moneys ...Money) Portfolio {
	for _, money := range moneys {
		p = append(p, money)
	}
	return p
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, m := range p {
		total += m.amount
	}
	return Money{amount: total, currency: currency}
}
