class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        best = 0
        nums = set(nums)

        for num in nums:
            if num - 1 not in nums:
                current = 1
                while num + current in nums:
                    current += 1
                best = max(current, best)

        
        return best
        
        

        




