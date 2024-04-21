package co

// HasTwoSum searches for sums by checking for the complementary difference with - values in the array and target
func HasTwoSum(nums []int, target int) bool {
	diffMap := make(map[int]bool)

	for _, num := range nums {
		if diffMap[num] {
			return true
		}
		
		diffMap[target-num] = true
	}

	return false
}