func punishmentNumber(n int) int {
    var output int
    
    for i:=1; i<=n; i++ {
        square:=i*i
        if checkValid(i, square) {
            output += square
        }
    }
    
    
    return output
}

func sum(data []int) int {
    var output int
    for _,v:=range data {
        output += v
    }
    return output
}

func dfs(square []string, index int, temp []int, target int) bool {
    if index == len(square) && sum(temp) == target {
        return true
    }

    if index >= len(square) || sum(temp) > target {
        return false
    }

    for i:=index; i<len(square); i++ {
        val,_ := strconv.Atoi(strings.Join(square[index:i+1], ""))
        temp = append(temp, val)
        if dfs(square, i+1, temp, target) {
            return true
        }
        temp = temp[:len(temp)-1]
    }

    return false
}

func checkValid(num, square int) bool {
    tempSquare := strings.Split(fmt.Sprintf("%d", square), "")
    return dfs(tempSquare, 0, []int{}, num)
}