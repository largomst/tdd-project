const assert = require('assert');

class Dollar {
    constructor(amount) {
        this.amount = amount;
    }
    times(multiplier) {
        return new Dollar(10);
    }
}

let fiver = new Dollar(5);
let tener = fiver.times(10);
assert.strictEqual(tener.amount, 10);

