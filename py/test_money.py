import unittest

from money import Money
from portfolio import Portfolio

class TestMoney(unittest.TestCase):
    def testMultiplicationInDollar(self):
        fiveDollars = Money(5, "USD")
        tenDollars = fiveDollars.times(2)
        self.assertEqual(tenDollars, fiveDollars.times(2))

    def testMultiplicationInEuros(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = tenEuros.times(2)
        self.assertEqual(twentyEuros, tenEuros.times(2))

    def testDivision(self):
        originalMoney = Money(4002, "KRW")
        actualMoneyAfterDivision = originalMoney.divide(4)
        expectedMoneyAfterDivision = Money(1000.5, "KRW")
        self.assertEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision)

    def testAddition(self):
        fiveDollars = Money(5, "USD")
        tenDollars = Money(10, "USD")
        fifteenDollars = Money(15, "USD")
        portfolio = Portfolio()
        portfolio.add(fiveDollars, tenDollars)
        self.assertEqual(portfolio.evaluate("USD"), fifteenDollars)


if __name__ == "__main__":
    unittest.main()
