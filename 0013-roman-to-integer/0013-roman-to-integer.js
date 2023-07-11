/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
    const symbol = {
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
    }

    let output = 0

    for (i = 0; i < s.length; i++) {
        const curr = symbol[s[i]]
        const next = symbol[s[i+1]]

        if (curr < next) {
            output += (next - curr)
            i++
        } else {
            output += curr
        }
    }

    return output
};