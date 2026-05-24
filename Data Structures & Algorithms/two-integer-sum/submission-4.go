func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := numsMap[diff]; ok {
			return []int{v, i}
		}
		numsMap[nums[i]] = i
	}
	return []int{}
}
