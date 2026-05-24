func hasDuplicate(nums []int) bool {
    numsMap := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if _, ok := numsMap[nums[i]]; ok {
            return true
        }
        numsMap[nums[i]] = true
    }
    return false
}
