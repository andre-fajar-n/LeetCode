/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    let xStr = x.toString()
    let newStr = xStr.split('').reverse().join('')
    let newNum = parseInt(newStr)
    
    if ( newNum > 0x7FFFFFFF) return 0

    return x < 0 ? -newNum : newNum
};