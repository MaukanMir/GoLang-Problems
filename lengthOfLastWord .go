

// 58. Length of Last Word
// Easy

// 259

// 33

// Add to List

// Share
// Given a string s consisting of some words separated by some number of spaces, return the length of the last word in the string.

// A word is a maximal substring consisting of non-space characters only.

 

// Example 1:

// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
// Example 2:

// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.
// Example 3:

// Input: s = "luffy is still joyboy"
// Output: 6
// Explanation: The last word is "joyboy" with length 6.
 

// Constraints:

// 1 <= s.length <= 104
// s consists of only English letters and spaces ' '.
// There will be at least one word in s.





// Success
// Details 
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
// Memory Usage: 2.2 MB, less than 86.55% of Go online submissions for Length of Last Word.




import ("strings")
func lengthOfLastWord(s string) int {
    
    transformed := strings.TrimSpace(s)
    count:=0
    for i:= len(transformed)-1; i>=0; i--{
        if string(transformed[i]) ==" "{
            return count
        }
        count+=1
    }
    
    return len(transformed)
}
