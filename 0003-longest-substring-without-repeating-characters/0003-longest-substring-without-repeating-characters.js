/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    let substring = ""
    let output = 0
    for (const idx in s) {
        const char = s[idx]
        if (substring.includes(char)) {
            console.log("CURR", idx, char)
            if (substring.length > output) {
                output = substring.length
            }
            console.log("SUB", substring.slice(substring.indexOf(char)+1))
            console.log("THIS", substring.indexOf(char))
            substring = substring.slice(substring.indexOf(char)+1)
        }
        substring += char
    }

    console.log("CEK", substring)
    if (substring.length > output) {
        output = substring.length
    }

    return output
};