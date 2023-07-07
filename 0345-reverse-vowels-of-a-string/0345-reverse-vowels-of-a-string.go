func reverseVowels(s string) string {
    vowels := "aiueoAIUEO"

    tempVowels := []struct{
        idx int
        char string
    } {}

    for i, v :=range s {
        if strings.Contains(vowels, string(v)) {
            tempVowels = append(tempVowels, struct{
                idx int
                char string
            }{
                idx: i,
                char: string(v),
            })
        }
    }

    iTempVowels := 0
    for i := len(s) - 1; i >= 0; i-- {
        char := string(s[i])
        if strings.Contains(vowels, string(char)) {
            s = s[:i] + tempVowels[iTempVowels].char + s[i+1:]
            iTempVowels++
        }
    }

    return s
}