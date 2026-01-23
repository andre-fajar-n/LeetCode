func minimumPairRemoval(nums []int) int {
    output:=0

    for len(nums) > 1 {
        new_nums, already_sorted := decreaseMinimumPair(nums)
        
        if already_sorted {
            break
        }
        output++
        nums = new_nums
    }

    return output
}

func decreaseMinimumPair(nums []int) ([]int, bool) {
    already_sorted := true
    minimum_sum_adjacent := math.MaxInt
    index_first_pair_adjacent := -1
    for i:=0; i<len(nums) - 1; i++ {
        minimum_sum := nums[i] + nums[i+1]
        if nums[i] > nums[i+1] {
            already_sorted = false
        }
        if minimum_sum < minimum_sum_adjacent {
            minimum_sum_adjacent = minimum_sum
            index_first_pair_adjacent = i
        }
    }

    if already_sorted {
        return nums, already_sorted
    }
    new_nums := append(nums[:index_first_pair_adjacent], minimum_sum_adjacent)
    new_nums = append(new_nums, nums[index_first_pair_adjacent+2:]...)

    return new_nums, already_sorted
}