// 9. Palindrome Number
// Easy

// 4255

// 1895

// Add to List

// Share
// Given an integer x, return true if x is palindrome integer.

// An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.

 

// Example 1:

// Input: x = 121
// Output: true
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
// Example 4:

// Input: x = -101
// Output: false

// Runtime: 28 ms, faster than 37.08% of Go online submissions for Palindrome Number.
// Memory Usage: 6.3 MB, less than 16.28% of Go online submissions for Palindrome Number.


import "strconv"
func isPalindrome(x int) bool {
    transform:= strconv.Itoa(x)
    ans :=""
    for i:= len(transform)-1;i>=0;i--{
        ans += string(transform[i])
    }
    
    check,_ := strconv.Atoi(ans)
    
    return check == x
}

