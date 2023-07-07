class Solution:
    def distributeCandies(self, candies: List[int]) -> int:
        if len(set(candies)) > len(candies)/2:
            return len(candies)//2
        else:
            return len(set(candies))