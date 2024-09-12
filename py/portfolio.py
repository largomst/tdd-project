import functools
import operator
from money import Money


class Portfolio:
    def __init__(self):
        self.moneys = []

    def add(self, *moneys):
        self.moneys += moneys

    def evaluate(self, bank, currency):
        total = 0.0
        failures = []
        for m in self.moneys:
            try:
                total += bank.convert(m, currency).amount
            except Exception as ex:
                failures.append(ex)
        if len(failures) == 0:
            return Money(total, currency)
        else:
            failureMessage = ",".join(str(f.args[0]) for f in failures)
            raise Exception(f"Missing exchange rate(s):[{failureMessage}]")
