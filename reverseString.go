// 344. Reverse String
// Easy

// 3280

// 845

// Add to List

// Share
// Write a function that reverses a string. The input string is given as an array of characters s.

 

// Example 1:

// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]
// Example 2:

// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]
 

// Constraints:

// 1 <= s.length <= 105
// s[i] is a printable ascii character.
 

Follow up: Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.



Success
Details 
Runtime: 28 ms, faster than 99.08% of Go online submissions for Reverse String.
Memory Usage: 6.7 MB, less than 87.21% of Go online submissions for Reverse String.



func reverseString(s []byte)  {
    left,right := 0, len(s)-1
    
    for left < right{
        s[left], s[right] = s[right],s[left]
        left+=1
        right -=1
    }
}
