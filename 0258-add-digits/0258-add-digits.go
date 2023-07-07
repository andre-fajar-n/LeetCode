func addDigits(num int) int {
    if num < 10 {
        return num
    }
    
    output := 0

    for num > 9 {
        output += num % 10
        num /= 10
    }

    output += num

    return addDigits(output)
}