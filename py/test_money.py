import unittest

from money import Money
from portfolio import Portfolio
from bank import Bank


class TestMoney(unittest.TestCase):
    def setUp(self):
        self.bank = Bank()
        self.bank.addExchangeRate("EUR", "USD", 1.2)
        self.bank.addExchangeRate("USD", "KRW", 1100)

    def testMultiplication(self):
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
        self.assertEqual(portfolio.evaluate(self.bank, "USD"), fifteenDollars)

    def testAdditionDollarAndEuro(self):
        fiveDollars = Money(5, "USD")
        tenEuros = Money(10, "EUR")
        portfolio = Portfolio()
        portfolio.add(fiveDollars, tenEuros)
        expectedResult = Money(17, "USD")
        actualResult = portfolio.evaluate(self.bank, "USD")
        self.assertEqual(expectedResult, actualResult)

    def testAdditionOfDollarsAndWons(self):
        oneDollar = Money(1, "USD")
        elevenHundredWon = Money(1100, "KRW")
        portfolio = Portfolio()
        portfolio.add(oneDollar, elevenHundredWon)
        expectedResult = Money(2200, "KRW")
        actualValue = portfolio.evaluate(self.bank, "KRW")
        self.assertEqual(
            expectedResult, actualValue, "%s != %s" % (expectedResult, actualValue)
        )

    def testAdditionWithMultipleMissingExchangeRates(self):
        oneDollar = Money(1, "USD")
        oneEuro = Money(1, "EUR")
        oneWon = Money(1, "KRW")
        portofolio = Portfolio()
        portofolio.add(oneDollar, oneEuro, oneWon)
        with self.assertRaisesRegex(
            Exception,
            r"Missing exchange rate\(s\):\[USD->Kalganid,EUR->Kalganid,KRW->Kalganid\]",
        ):
            portofolio.evaluate(self.bank, "Kalganid")

    def testConversionWithDifferentRatesBetweenTwoCurrencies(self):
        tenEuros = Money(10, "EUR")
        self.assertEqual(self.bank.convert(tenEuros, "USD"), Money(12, "USD"))

        self.bank.addExchangeRate("EUR", "USD", 1.3)
        self.assertEqual(self.bank.convert(tenEuros, "USD"), Money(13, "USD"))

    def testConversionWithMissingExchangeRate(self):
        bank = Bank()
        tenEuros = Money(10, "EUR")
        with self.assertRaisesRegex(Exception, r"EUR->Kalganid"):
            bank.convert(tenEuros, "Kalganid")
    
    def testWhatIsTheConversionRateRomEURtoUSD(self):
        tenEuros = Money(10, "EUR")
        self.assertEqual(self.bank.convert(tenEuros, "USD"), Money(12, "USD"))


if __name__ == "__main__":
    unittest.main()
