/**
 * @param {number} num
 * @return {boolean}
 */
var checkPerfectNumber = function(num) {
    let num1 = 1
    let output = 0
    while (num1 <= Math.sqrt(num)) {
        if (num % num1 == 0) {
            let num2 = num / num1
            if (num2 == num) {
                num2 = 0
            }
            output = output + num1 + num2
        }
        num1++
    }
    
    return output == num && num != 1
};