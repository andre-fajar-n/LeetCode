func minimumAbsDifference(arr []int) [][]int {
    max_int := math.MaxInt
    output := make([][]int, 0)
    sort.Ints(arr)
    for i := 0; i < len(arr) - 1; i++ {
        diff := arr[i+1] - arr[i]
        greater_num := arr[i+1]
        lower_num := arr[i]
        diff_abs := int(math.Abs(float64(diff)))
        if diff_abs < max_int {
            max_int = diff_abs
            output = [][]int{{lower_num, greater_num}}
        } else if diff_abs == max_int {
            output = append(output, []int{lower_num, greater_num})
        }
    }

    return output
}