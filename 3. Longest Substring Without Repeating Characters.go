// Given a string s, find the length of the longest substring without repeating characters.

 

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 

// Constraints:

// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.

// Success
// Details 
// Runtime: 12 ms, faster than 48.14% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 3.4 MB, less than 45.83% of Go online submissions for Longest Substring Without Repeating Characters.


func lengthOfLongestSubstring(s string) int {
    charSet := make(map[string]bool)
    
    var leftPointer =0
    var result = float64(0)
    
    for rightPointer:=0; rightPointer< len(s); rightPointer ++{
        for(charSet[string(s[rightPointer])] == true){
            delete(charSet,string(s[leftPointer]))
            leftPointer +=1
        }
        
        charSet[string(s[rightPointer])] = true
        result = math.Max(float64(result), float64(rightPointer) - float64(leftPointer -1))
    }
    
    return int(result)
}
