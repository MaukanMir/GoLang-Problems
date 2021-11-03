// 383. Ransom Note
// Easy

// 1255

// 269

// Add to List

// Share
// Given two stings ransomNote and magazine, return true if ransomNote can be constructed from magazine and false otherwise.

// Each letter in magazine can only be used once in ransomNote.

 

// Example 1:

// Input: ransomNote = "a", magazine = "b"
// Output: false
// Example 2:

// Input: ransomNote = "aa", magazine = "ab"
// Output: false
// Example 3:

// Input: ransomNote = "aa", magazine = "aab"
// Output: true




import "strings"
func canConstruct(ransomNote string, magazine string) bool {
    hashmap := make(map[string]int,0)
    ans:=""
    
    for i:= range magazine{
        check := string(magazine[i])
        
         if hashmap[check] ==0{
            hashmap[check]=1
        }else{
            hashmap[check] +=1
        }
        
        
    }
    
    for i:= range ransomNote{
        check := string(ransomNote[i])
        
        if strings.Contains(magazine,check) == false{
            return false
        }else if strings.Contains(magazine,check) && hashmap[check] >0{
            hashmap[check] -=1
            ans += check
            
        }
    }
    
    return ransomNote == ans
}
