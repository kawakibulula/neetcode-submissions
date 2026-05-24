func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]
        if num, ok := numsMap[complement]; ok {
            return []int{num, i}
        }
        numsMap[nums[i]] = i
    }
    return []int{}
}
