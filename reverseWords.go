
// 557. Reverse Words in a String III
// Easy

// 1995

// 128

// Add to List

// Share
// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

 

// Example 1:

// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
// Example 2:

// Input: s = "God Ding"
// Output: "doG gniD"
 

// Constraints:

// 1 <= s.length <= 5 * 104
// s contains printable ASCII characters.
// s does not contain any leading or trailing spaces.
// There is at least one word in s.
// All the words in s are separated by a single space.





import("strings")
func reverseWords(s string) string {
    words := strings.Fields(s)
    ans :=""
    
    for _, word:= range words{
        for i:= len(word)-1; i>=0;i--{
            ans += string(word[i])
        }
        ans +=  " "
    }
    
    return strings.TrimSpace(ans)
}




Success
Details 
Runtime: 4 ms, faster than 94.59% of Go online submissions for Reverse Words in a String III.
Memory Usage: 6.5 MB, less than 92.62% of Go online submissions for Reverse Words in a String III.



import("strings")
func reverseString(s []byte)  {
    length := len(s)
    for i := 0; i < length / 2; i++ {
      s[i], s[length-i-1] = s[length-i-1], s[i]
    }
}

func reverseWords(s string) string {
    words := strings.Split(s, " ")
    for k, v := range words {
        rev := []byte(v)
        reverseString(rev)
        words[k] = string(rev)
    }
    return strings.Join(words, " ")
}
