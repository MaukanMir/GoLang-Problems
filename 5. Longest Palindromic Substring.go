// Given a string s, return the longest palindromic substring in s.

 

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"
 

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters.

// Success
// Details 
// Runtime: 4 ms, faster than 95.65% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Longest Palindromic Substring.



func longestPalindrome(s string) string {
    
    
    var res = ""
    var resLen =0
    
    for i:=0; i< len(s); i++{
        
        var L =i
        var R =i
        
        for L>=0 && R< len(s) && s[L] ==s[R]{
            if(R-L +1 > resLen){
                resLen = R-L +1
                res = s[L:R+1]
            }
            L-=1
            R +=1
        }
        
        
        L =i
        R = i+1
        
        for L >=0 && R< len(s) && s[L] == s[R]{
            if(R-L +1 > resLen){
                resLen = R-L +1
                res = s[L:R+1]
            }
            L-=1
            R +=1
        }
    }
    
    return res
    
    
    
    
}
