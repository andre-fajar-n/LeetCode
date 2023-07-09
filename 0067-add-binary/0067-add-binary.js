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
        console.log("AWAL", output, "CARRY", carry)
        if (i >= 0) {
            carry += parseInt(a[i])
            console.log("A", a[i], carry)
            i--
        }

        if (j >= 0) {
            carry += parseInt(b[j])
            console.log("B", b[j], carry)
            j--
        }

        const temp = carry % 2
        console.log("TEMP", temp, carry, output)
        output = temp.toString() + output
        carry = parseInt(carry/2)
        console.log("AKHIR", output, "CARRY", carry)
    }

    return output
};