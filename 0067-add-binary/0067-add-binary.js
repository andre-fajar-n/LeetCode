/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function(a, b) {
    let i = a.length - 1
    let j = b.length - 1
    let carry = 0
    let output = ""

    while (i >= 0 || j >= 0 || carry == 1) {
        if (i >= 0) {
            carry += parseInt(a[i])
            i--
        }

        if (j >= 0) {
            carry += parseInt(b[j])
            j--
        }

        const temp = carry % 2
        output = temp.toString() + output
        carry = parseInt(carry/2)
    }

    return output
};