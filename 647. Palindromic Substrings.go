Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.

 

Example 1:

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
Example 2:

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".



Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Palindromic Substrings.
Memory Usage: 1.8 MB, less than 97.39% of Go online submissions for Palindromic Substrings.

func countSubstrings(s string) int {
    
    var result int
    
    for i:=0; i< len(s); i++{
        result += findPalis(s,i,i)
        result += findPalis(s,i, i+1)
    }
    
    return result
}

func findPalis(str string, lo int, high int)int{
    var result int
    
    for lo>=0 && high < len(str){
        if str[lo] != str[high]{break}
        lo --
        high++
        result ++
    }
    
    return result
}
