
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
