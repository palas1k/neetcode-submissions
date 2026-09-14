func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range(nums) {
		freqMap[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range freqMap {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i>= 0; i-- {
		if len(buckets[i]) > 0 {
			for _, num := range buckets[i] {
				result = append(result, num)

				if len(result) == k {
					return result
				}
			}
		}
	}
	return result
}
