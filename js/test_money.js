const assert = require('assert');

class Dollar {
    constructor(amount) {
        this.amount = amount;
    }
    times(multiplier) {
        return new Dollar(this.amount * multiplier);
    }
}

let fiver = new Dollar(5);
let tener = fiver.times(2);
assert.strictEqual(tener.amount, 10);

