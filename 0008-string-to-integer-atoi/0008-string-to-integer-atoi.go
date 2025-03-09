func myAtoi(s string) int {
    outputStr:=""
    s=strings.TrimSpace(s)
    for i:=0; i<len(s); i++ {
        single:=string(s[i])
        _, err := strconv.Atoi(single)
        if err != nil && (single != "-"  && single != "+") {
            break
        }
        if (single == "-" || single == "+") && outputStr != "" {
            break
        }
        outputStr += single
    }
    output,_:=strconv.Atoi(outputStr)
    if output > 2147483647 {
        output = 2147483647
    }
    if output < -2147483648 {
        output = -2147483648
    }
    return output
}