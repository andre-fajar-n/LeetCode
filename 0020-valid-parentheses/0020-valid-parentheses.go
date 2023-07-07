func isValid(s string) bool {
    brackets := []string{}
    closeBrackets := []string{}

    for _, v := range s {
        char := string(v)
        if char == "(" || char == "{" || char == "[" {
            brackets = append(brackets, char)
        } else {
            lastChar := ""
            if len(brackets) > 0 {
                lastChar = brackets[len(brackets) - 1]
            }
            haveOpenBracket := false
            switch char {
                case ")":
                    if lastChar == "(" {
                        haveOpenBracket = true
                    }
                case "}":
                    if lastChar == "{" {
                        haveOpenBracket = true
                    }
                case "]":
                    if lastChar == "[" {
                        haveOpenBracket = true
                    }
            }

            if haveOpenBracket {
                brackets = brackets[:len(brackets) - 1]
            } else {
                closeBrackets = append(closeBrackets, char)
            }
        }
    }

    return len(closeBrackets) == 0 && len(brackets) == 0
}