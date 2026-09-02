func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		seek := target-nums[i]
		if idx, exists := seen[seek]; exists {
			return []int{idx, i}
		}
		seen[nums[i]] = i
	}
	return nil
}

