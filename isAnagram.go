




// 242. Valid Anagram
// Easy

// 3496

// 188

// Add to List

// Share
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

 

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:

// Input: s = "rat", t = "car"
// Output: false
 

// Constraints:

// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.



// Success
// Details 
// Runtime: 12 ms, faster than 31.98% of Go online submissions for Valid Anagram.
// Memory Usage: 3.7 MB, less than 19.16% of Go online submissions for Valid Anagram.





func isAnagram(s string, t string) bool {
    one:= make(map[string]int)
    two:= make(map[string]int)
    
    for i:= range s{
        one[string(s[i])]++
    }
    
    for i:= range t{
        two[string(t[i])]++
    }
    
    for i:= range one{
        
        if one[i] > two[i]{return false}
    }
    
    for i:= range two{
        
        if two[i] > one[i]{return false}
    }
    
    return true
}
