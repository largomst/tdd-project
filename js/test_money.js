const assert = require('assert');

class Money {
    constructor(amount, currency) {
        this.amount = amount;
        this.currency = currency;
    }
    times(multiplier) {
        return new Money(this.amount * multiplier, this.currency);
    }
    divide(divisor) {
        return new Money(this.amount / divisor, this.currency);
    }
}

let fiver = new Money(5, "USD");
assert.strictEqual((fiver.times(2)).amount, 10);

let tenEuros = new Money(10, "EUR");
assert.strictEqual((tenEuros.times(2)).amount, 20);
assert.strictEqual((tenEuros.times(2)).currency, "EUR");

let originalMoney = new Money(4002, "KRW");
let actualMoneyAfterDivision = originalMoney.divide(4);
let expectedMoneyAfterDivision = new Money(1000.5, "KRW");
assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision);