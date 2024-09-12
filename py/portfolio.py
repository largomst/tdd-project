import functools
import operator
from money import Money


class Portfolio:
    def __init__(self):
        self.moneys = []
        self._eur_to_usd = 1.2

    def add(self, *moneys):
        self.moneys += moneys

    def evaluate(self, currency):
        total = functools.reduce(operator.add, map(lambda m: self._convert(m, currency), self.moneys), 0)
        return Money(total, currency)

    def _convert(self, aMoney, currency):
        if aMoney.currency == currency:
            return aMoney.amount
        else:
            return aMoney.amount * self._eur_to_usd

