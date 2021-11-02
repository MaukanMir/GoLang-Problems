// 387. First Unique Character in a String
// Easy

// 3844

// 169

// Add to List

// Share
// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

 

// Example 1:

// Input: s = "leetcode"
// Output: 0
// Example 2:

// Input: s = "loveleetcode"
// Output: 2
// Example 3:

// Input: s = "aabb"
// Output: -1
 

// Constraints:

// 1 <= s.length <= 105
// s consists of only lowercase English letters.




// Runtime: 4 ms, faster than 97.96% of Go online submissions for First Unique Character in a String.
// Memory Usage: 5.4 MB, less than 100.00% of Go online submissions for First Unique Character in a String.





import("strings")
func firstUniqChar(s string) int {
    

    for i:=0;i<len(s);i++{
        got:= string(s[i])
        
        if strings.Contains(s[i+1:],got) == false && strings.Contains(s[:i],got) ==false{
            return i
        }
    
    
}
    
     return -1
    
}
