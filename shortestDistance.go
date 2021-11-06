// 243. Shortest Word Distance
// Easy

// 859

// 72

// Add to List

// Share
// Given an array of strings wordsDict and two different strings that already exist in the array word1 and word2, return the shortest distance between these two words in the list.

 

// Example 1:

// Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "coding", word2 = "practice"
// Output: 3
// Example 2:

// Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "makes", word2 = "coding"
// Output: 1
 

// Constraints:

// 1 <= wordsDict.length <= 3 * 104
// 1 <= wordsDict[i].length <= 10
// wordsDict[i] consists of lowercase English letters.
// word1 and word2 are in wordsDict.
// word1 != word2




Success
Details 
Runtime: 4 ms, faster than 96.43% of Go online submissions for Shortest Word Distance.
Memory Usage: 4.5 MB, less than 37.50% of Go online submissions for Shortest Word Distance.




import "math"

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    indexOne, indexTwo := float64(-1),float64(-1)
    min_distance := float64(len(wordsDict))
    
    for i:= range wordsDict{
        if string(wordsDict[i])  == word1{
            indexOne = float64(i)
        }else if string(wordsDict[i]) == word2{
            indexTwo = float64(i)
        }
        
        if indexOne != -1 && indexTwo != -1{
            min_distance = math.Min(float64(min_distance), math.Abs(float64(indexOne) - float64(indexTwo)))
        }
    }
    
    return int(min_distance)
}
