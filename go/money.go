package main

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
