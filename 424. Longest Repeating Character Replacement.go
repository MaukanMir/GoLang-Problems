// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can get after performing the above operations.

 

// Example 1:

// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.
// Example 2:

// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.

// Success
// Details 
// Runtime: 16 ms, faster than 33.83% of Go online submissions for Longest Repeating Character Replacement.
// Memory Usage: 3.5 MB, less than 6.77% of Go online submissions for Longest Repeating Character Replacement.


func characterReplacement(s string, k int) int {
    var maxCount = 0;
    var leftPointer = 0;
    var result = 0;
    count := make(map[string]int)
    
    for rightPointer :=0; rightPointer < len(s); rightPointer ++{
        
        if count[string(s[rightPointer])] == 0{ count[string(s[rightPointer])] =1
        }else {count[string(s[rightPointer])] += 1}
        
        maxCount = int(math.Max(float64(maxCount), float64(count[string(s[rightPointer])])))
        
        for (rightPointer - leftPointer +1) - maxCount > k{
            count[string(s[leftPointer])] -=1
            leftPointer +=1
        }
        
        result = int(math.Max(float64(result), float64(rightPointer -leftPointer +1)))
    }
    
    return result
}



func characterReplacement(s string, k int) int {
    var maxCount = 0;
    var leftPointer = 0;
    var result = 0;
    count := make(map[string]int)
    
    for rightPointer :=0; rightPointer < len(s); rightPointer ++{
        
        if count[string(s[rightPointer])] == 0{ count[string(s[rightPointer])] =1
        }else {count[string(s[rightPointer])] += 1}
        
        maxCount = max(maxCount, count[string(s[rightPointer])])
        
        for (rightPointer - leftPointer +1) - maxCount > k{
            count[string(s[leftPointer])] -=1
            leftPointer +=1
        }
        
        result = max(result, rightPointer -leftPointer +1)
    }
    
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
