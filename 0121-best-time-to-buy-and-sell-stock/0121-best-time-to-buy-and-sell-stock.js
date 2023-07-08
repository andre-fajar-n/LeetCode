/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let output = 0
    let buyPrice = prices[0]
    for (let i = 1; i < prices.length; i++) {
        const price = prices[i]
        if (price < buyPrice) {
            buyPrice = price
        } else if (price > buyPrice) {
            let temp = price - buyPrice
            if (temp > output) {
                output = temp
            }
        }
    }

    return output
};