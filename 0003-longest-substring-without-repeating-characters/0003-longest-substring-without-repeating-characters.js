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
            if (substring.length > output) {
                output = substring.length
            }
            substring = substring.slice(substring.indexOf(char)+1)
        }
        substring += char
    }

    if (substring.length > output) {
        output = substring.length
    }

    return output
};