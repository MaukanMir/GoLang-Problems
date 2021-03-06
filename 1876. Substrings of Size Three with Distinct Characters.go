// A string is good if there are no repeated characters.

// Given a string s​​​​​, return the number of good substrings of length three in s​​​​​​.

// Note that if there are multiple occurrences of the same substring, every occurrence should be counted.

// A substring is a contiguous sequence of characters in a string.

 

// Example 1:

// Input: s = "xyzzaz"
// Output: 1
// Explanation: There are 4 substrings of size 3: "xyz", "yzz", "zza", and "zaz". 
// The only good substring of length 3 is "xyz".
// Example 2:

// Input: s = "aababcabc"
// Output: 4
// Explanation: There are 7 substrings of size 3: "aab", "aba", "bab", "abc", "bca", "cab", and "abc".
// The good substrings are "abc", "bca", "cab", and "abc".


Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Substrings of Size Three with Distinct Characters.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Substrings of Size Three with Distinct Characters.



func countGoodSubstrings(s string) int {
    start , end := 0, len(s)
    
    if end < 3{return 0}
    
    for i:= 0; i< end-2; i++{
        checkValues,check := [26]int{}, true
        
        for j := i; j< i+3; j++{
            checkValues[s[j] -'a']++
            if checkValues[s[j] -'a'] > 1 {check = false; break}
        }
        if check{start ++}
    }
    
    return start
    
}
