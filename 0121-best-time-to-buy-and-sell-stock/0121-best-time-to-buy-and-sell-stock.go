func maxProfit(prices []int) int {
    output := 0
    buyPrice := prices[0]
    for i:= 1; i < len(prices); i++ {
        price := prices[i]
        if price < buyPrice {
            buyPrice = price
        } else if price > buyPrice {
            temp := price - buyPrice
            if temp > output {
                output = temp
            }
        }
    }

    return output
}