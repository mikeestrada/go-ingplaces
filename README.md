# Go-ing places (hopefully)

A repo of practice and actual interview questions.

Expected output/results:
```
----fizz buzz----: 
[1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizz 16 17 fizz 19 buzz]


----palidrome----: 
false
true

// HasTwoSum searches for sums by checking for the complementary difference with - values in the array and target
----Has Two Sums----: 
true

// PinpointZero first searches for zeros and takes note of their address in the matrix.
// Later iterates through each row/col to zero out the flagged rows/cols.
----pin point zeros in matrix----: 
[[1 0 3 0] [0 0 0 0] [9 0 11 0] [0 0 0 0]]

// WIP
----sum linked lists----: 
<nil>

// ReverseWords uses a tuple swap assignment to reverse. 2 pointers operate against a slice of words.
// Tuple assignments are atomic in nature and therefore avoid interference.
----reverse words----: 
me reverse
----reverse strin----: 
em esrever

// GetPalindromes - finds palindrome substrings within a given string of words.
// Uses 2 pointers: first pointer is root and second denotes the end of the substring. Each letter is root at least once.
// Second pointer starts at i+2 because a palindrome must be at min 3 letters.
----get palindromes in string----: 
[rever eve racecar aceca cec]

// AggregateReceivables - aggregates receivables by merchant, card type. (Assumption) Date in result is last transaction date.
----aggregate Receivables----: 
merchant,cardType,txDate,amount
merchA,mc,2020-01-03,299
merchA,visa,2020-01-03,349
merchB,mc,2020-01-02,10
merchB,visa,2020-01-02,20
```