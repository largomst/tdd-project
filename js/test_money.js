const assert = require('assert');

let fiver = new Dollor(5);
let tener = new fiver.timer(10);
assert.strictEqual(tener.amount, 10);