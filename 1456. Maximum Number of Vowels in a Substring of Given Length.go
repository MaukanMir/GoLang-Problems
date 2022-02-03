Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

 

Example 1:

Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.
Example 2:

Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.
Example 3:

Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.

Success
Details 
Runtime: 20 ms, faster than 17.24% of Go online submissions for Maximum Number of Vowels in a Substring of Given Length.
Memory Usage: 4.7 MB, less than 96.55% of Go online submissions for Maximum Number of Vowels in a Substring


func maxVowels(s string, k int) int {
    
    vowels := "aeiou"
    var maxAmount, count int
   
    
    for i:=0; i< k; i++{
        if strings.Contains(vowels,string(s[i])){count++}
    }
    
    maxAmount = count
    for i:=0; i< len(s)-k; i++{
        if strings.Contains(vowels,string(s[i])){count--}
        
        if strings.Contains(vowels,string(s[i+k])){count++}
        
        maxAmount = max(maxAmount,count)
    }
    
    return maxAmount
}

func max(a,b int)int{
    if a>b{
        return a
    }
    
    return b
}
