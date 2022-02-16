

Your friend is typing his name into a keyboard. Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard. Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.

 

Example 1:

Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.
Example 2:

Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it was not in the typed output.


Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Long Pressed Name.
Memory Usage: 2 MB, less than 93.75% of Go online submissions for Long Pressed Name.


func isLongPressedName(name string, typed string) bool {
    
    nameL, typedL := 0,0
    
    for nameL < len(name) && typedL < len(typed){
        
        if name[nameL] == typed[typedL]{
            nameL ++
            typedL ++
        }else if nameL > 0 && name[nameL-1] == typed[typedL]{
            typedL++
        }else{
            return false
        }   
    }
    
    for typedL < len(typed){
        if name[nameL-1] != typed[typedL]{
            return false
        }else{
            typedL++
        }
    }
    
    if nameL < len(name){return false}
    
    return true
}
