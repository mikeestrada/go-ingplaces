package co

import "strings"

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

type Node struct {
	Val int
	Next *Node
}

func SumLinkedLists(l1, l2 *Node) *Node {
	summed := &Node{}
	current := summed
	carryOver := 0

	for l1 != nil || l2 != nil || carryOver > 0 {
		sum := carryOver

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carryOver = sum / 10
		current.Next = &Node{Val: sum % 10}
		current = current.Next
	}

	return current.Next
}

func ReverseWords (s []byte) string {
	// convert byte slice to string
	str := string(s)

	// split string on spaces
	words := strings.Fields(str)

	// swap
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// create resulting string from reversed slice
	reversed := strings.Join(words, " ")

	return reversed
}