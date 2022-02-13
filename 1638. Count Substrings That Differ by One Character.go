Given two strings s and t, find the number of ways you can choose a non-empty substring of s and replace a single character by a different character such that the resulting substring is a substring of t. In other words, find the number of substrings in s that differ from some substring in t by exactly one character.

For example, the underlined substrings in "computer" and "computation" only differ by the 'e'/'a', so this is a valid way.

Return the number of substrings that satisfy the condition above.

A substring is a contiguous sequence of characters within a string.

 

Example 1:

Input: s = "aba", t = "baba"
Output: 6
Explanation: The following are the pairs of substrings from s and t that differ by exactly 1 character:
("aba", "baba")
("aba", "baba")
("aba", "baba")
("aba", "baba")
("aba", "baba")
("aba", "baba")
The underlined portions are the substrings that are chosen from s and t.
​​Example 2:
Input: s = "ab", t = "bb"
Output: 3
Explanation: The following are the pairs of substrings from s and t that differ by 1 character:
("ab", "bb")
("ab", "bb")
("ab", "bb")
​​​​The underlined portions are the substrings that are chosen from s and t.
 

Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Count Substrings That Differ by One Character.
Memory Usage: 1.8 MB, less than 100.00% of Go online submissions for Count Substrings That Differ by One Character.


func countSubstrings(s string, t string) int {
    
    var count int
    
    for one :=0; one < len(s); one++{
        for two:=0 ; two < len(t); two++{
            x:= one
            y:=two
            var diff int
            
            for x <len(s) && y < len(t){
                if s[x] != t[y]{
                    diff +=1
                }
                if diff ==1{count++}
                if(diff ==2){break}
                x++
                y++
            }
        }
    }
    
    return count
}
