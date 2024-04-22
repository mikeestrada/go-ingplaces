package main

import (
	"fmt"

	"go-ing.com/places/co"
)

func main() {
	fmt.Println("----fizz buzz----: ")
	fmt.Println(fizzBuzz(20))

	fmt.Println("----palidrome----: ")
	fmt.Println(co.IsPalindrome("palidrome"))
	fmt.Println(co.IsPalindrome("racecar"))

	fmt.Println("----Has Two Sums----: ")
	fmt.Println(co.HasTwoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println("----pin point zeros in matrix----: ")
	fmt.Println(co.PinpointZero([][]int{
		{1, 2, 3, 4},
		{5, 0, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 0},
	}))

	fmt.Println("----sum linked lists----: ")
	l1 := &co.Node{Val: 2, Next: &co.Node{Val: 4, Next: &co.Node{Val: 3}}}
	l2 := &co.Node{Val: 5, Next: &co.Node{Val: 6, Next: &co.Node{Val: 4}}}
	fmt.Println(co.SumLinkedLists(l1, l2))

	fmt.Println("----reverse words----: ")
	fmt.Println(co.ReverseWords([]byte("reverse me")))

	fmt.Println("----get palindromes in string----: ")
	fmt.Println(co.GetPalindromes(("reverse me racecar")))
}

func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			result = append(result, "fizz")
		} else if i%5 == 0 {
			result = append(result, "buzz")
		} else if i%3 == 0 && i%5 == 0 {
			result = append(result, "fizzbuzz")
		} else {
			result = append(result, fmt.Sprintf("%v", i))
		}
	}
	return result
}
