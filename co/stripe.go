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

func PinpointZero(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	rowsWithZeros := make(map[int]bool, rows)
	colsWithZeros := make(map[int]bool, cols)

	// search for zeros and flag
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				rowsWithZeros[i] = true
				colsWithZeros[j] = true
			}
		}
	}

	// zero out rows
	for i := 0; i < rows; i++ {
		if rowsWithZeros[i] {
			for j := 0; j < cols; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// zero out cols
	for j := 0; j < cols; j++ {
		if colsWithZeros[j] {
			for i := 0; i < rows; i++ {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}
