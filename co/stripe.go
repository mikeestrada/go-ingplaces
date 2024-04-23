package co

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

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

// PinpointZero first searches for zeros and takes note of their address in the matrix.
// Later iterates through each row/col to zero out the flagged rows/cols.
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
	Val  int
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

// ReverseWords uses a tuple swap assignment to reverse. 2 pointers operate against a slice of words.
// Tuple assignments are atomic in nature and therefore avoid interference.
func ReverseWords(s []byte) string {
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

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// GetPalindromes - finds palindrome substrings within a given string of words.
// Uses 2 pointers: first pointer is root and second denotes the end of the substring. Each letter is root at least once.
// Second pointer starts at i+2 because a palindrome must be at min 3 letters.
func GetPalindromes(s string) []string {
	var result []string

	for i := 0; i < len(s); i++ {
		for j := i + 2; j <= len(s); j++ {
			word := s[i:j]
			if IsPalindrome(word) && len(word) > 1 {
				result = append(result, word)
			}
		}
	}

	return result
}

type receivable struct {
	txDate string
	amount int
}

// AggregateReceivables - aggregates receivables by merchant, card type. (Assumption) Date in result is last transaction date.
func AggregateReceivables(s string) string {
	lines := strings.Split(s, "\n")
	receivables := ``
	merchantTypeAggs := make(map[string]map[string]receivable)

	// iterate through each line of receivables excluding the header
	for idx, line := range lines {
		if idx != 0 {
			date := line[strings.LastIndex(line, ",")+1:]
			dateRemainder := line[:strings.LastIndex(line, ",")]

			amount := dateRemainder[strings.LastIndex(dateRemainder, ",")+1:]
			amountRemainder := dateRemainder[:strings.LastIndex(dateRemainder, ",")]

			cardType := amountRemainder[strings.LastIndex(amountRemainder, ",")+1:]
			cardTypeRemainder := amountRemainder[:strings.LastIndex(amountRemainder, ",")]

			merchant := cardTypeRemainder[strings.LastIndex(cardTypeRemainder, ",")+1:]

			amtInt, err := strconv.ParseInt(amount, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			if merchantTypeAggs[merchant] == nil {
				merchantTypeAggs[merchant] = make(map[string]receivable)
			}

			merchantTypeAggs[merchant][cardType] = receivable{
				txDate: date,
				amount: int(amtInt) + merchantTypeAggs[merchant][cardType].amount,
			}
		}
	}

	// generate string of aggregated receivables as expected
	for merchant, cardTypes := range merchantTypeAggs {
		for cardType, receivable := range cardTypes {
			receivables += fmt.Sprintf("%v,%v,%v,%v\n", merchant, cardType, receivable.txDate, receivable.amount)
		}
	}

	return receivables
}

// middle of 2 sorted arrays. resort by using 2 pointers and merging into new array.
// // return middle or avg of middle 2.
// func arrarysMedian(a1, a2 []string) string {
// 	if a1 == nil || a2 == nil {
// 		return "error"
// 	}
// 	pointer1 := 0
// 	pointer2 := 0

// 	combinedLen := len(a1) + len(a2)

// 	if combinedLen % 2 == 0 {
// 		return
// 	}
// }
