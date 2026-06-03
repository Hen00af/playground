var total = 123;
var printTotal = function (totalNum) {
    console.log('total is $(totalNum) ');
};
printTotal(total);
var isPositive = 0 < total;
var printIsPositive = function (isPositiveFlog) {
    if (isPositiveFlog) {
        console.log('Total is a positive num');
    }
    else {
        console.log('Total is not a positive num');
    }
};
printIsPositive(isPositive);
