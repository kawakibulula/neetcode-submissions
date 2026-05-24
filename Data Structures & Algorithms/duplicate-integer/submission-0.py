class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        duplicate = {}
        for i, val in enumerate(nums):
            if val in duplicate.values():
                return True
            duplicate[i] = val
        return False
         