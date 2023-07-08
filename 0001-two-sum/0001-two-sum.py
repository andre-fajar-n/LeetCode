class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        lst = []
        for i in range(len(nums)):
            for k in range(i+1, len(nums)):
                save = nums[i] + nums[k]
                if save == target:
                    lst.append(i)
                    lst.append(k)
                    return lst                    