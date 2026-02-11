func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for idx, val := range nums {
        hashMap[val] = idx
    }

    for idx, val := range nums {
        diff := target - val
        secondIdx, exist := hashMap[diff]
        if idx != secondIdx && exist {
            return []int{idx, secondIdx}
        }
    }

    return []int{0, 0}
}