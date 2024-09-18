package main

import (
	s "tdd/stocks"
	"testing"
)

func TestMultiplication(t *testing.T) {
	tenEuros := s.NewMoney(10, "EUR")
	actualResult := tenEuros.Times(2)
	expectedResult := s.NewMoney(20, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := s.NewMoney(4002, "KRW")
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := s.NewMoney(1000.5, "KRW")
	assertEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

func TestAddition(t *testing.T) {
	var portfolio s.Portfolio
	var portfolioDollars s.Money

	fiveDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "USD")
	fifteenDollars := s.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars, tenDollars)
	portfolioDollars, _ = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portfolioDollars)
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio s.Portfolio
	fiveDollars := s.NewMoney(5, "USD")
	tenEuros := s.NewMoney(10, "EUR")
	portfolio = portfolio.Add(fiveDollars, tenEuros)

	expectedResult := s.NewMoney(17, "USD")
	actualResult, _ := portfolio.Evaluate("USD")

	assertEqual(t, expectedResult, actualResult)

}

func TestAdditionDollarsAndWons(t *testing.T) {
	var portfolio s.Portfolio

	fiveDollars := s.NewMoney(1, "USD")
	elevenHundredDollars := s.NewMoney(1100, "KRW")
	portfolio = portfolio.Add(fiveDollars, elevenHundredDollars)

	expectedValue := s.NewMoney(2200, "KRW")
	actualValue, _ := portfolio.Evaluate("KRW")

	assertEqual(t, expectedValue, actualValue)
}

func TestAdditionWithMultipleMissExchangeRates(t *testing.T) {
	var portfolio s.Portfolio

	oneDollar := s.NewMoney(1, "USD")
	oneEuro := s.NewMoney(1, "EUR")
	oneWon := s.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar, oneEuro, oneWon)

	expectErrorMessage :=
		"Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	_, actualErro := portfolio.Evaluate("Kalganid")

	assertEqual(t, expectErrorMessage, actualErro.Error())
}
