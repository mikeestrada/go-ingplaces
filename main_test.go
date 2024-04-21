package main

import "testing"

func testIsPalindrome(t *testing.T) {
	t.Error(isPalindrome("racecar") == true)
	t.Error(isPalindrome("palindrome") == false)
}

func testFizzBuzz(t *testing.T) {
	t.Error(fizzBuzz(20)[0] != "1")
}