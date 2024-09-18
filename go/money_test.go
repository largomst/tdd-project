package main

import (
	s "tdd/stocks"
	"testing"
)

func TestMultiplicationDollar(t *testing.T) {
	five := s.Money{
		amount:   5,
		currency: "USD",
	}
	actualResult := five.Times(2)
	expectedResult := s.Money{
		amount:   10,
		currency: "USD",
	}
	assertEqual(t, expectedResult, actualResult)
}
func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := s.Money{amount: 10, currency: "EUR"}
	actualResult := tenEuros.Times(2)
	expectedResult := s.Money{amount: 20, currency: "EUR"}
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := s.Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := s.Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

func TestAddition(t *testing.T) {
	var portfolio s.Portfolio
	var portfolioDollars s.Money

	fiveDollars := s.Money{amount: 5, currency: "USD"}
	tenDollars := s.Money{amount: 10, currency: "USD"}
	fifteenDollars := s.Money{amount: 15, currency: "USD"}

	portfolio = portfolio.Add(fiveDollars, tenDollars)
	portfolioDollars = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portfolioDollars)
}

func assertEqual(t *testing.T, expected s.Money, actual s.Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}
