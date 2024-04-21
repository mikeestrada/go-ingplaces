package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("----fizz buzz----: ")
	fmt.Println(fizzBuzz(20))
	fmt.Println("----palidrome----: ")
	fmt.Println(isPalindrome("palidrome"))
	fmt.Println(isPalindrome("racecar"))

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

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
