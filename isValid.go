

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
 

// Example 1:

// Input: s = "()"
// Output: true
// Example 2:

// Input: s = "()[]{}"
// Output: true
// Example 3:

// Input: s = "(]"
// Output: false
// Example 4:

// Input: s = "([)]"
// Output: false
// Example 5:

// Input: s = "{[]}"
// Output: true














// Success
// Details 
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.2 MB, less than 23.96% of Go online submissions for Valid Parentheses.




func isValid(s string) bool {
    parens:= s
    
    opening := "([{"
    closing := "]})"

    dict := map[interface{}]interface{} {
        ")": "(",
        "}": "{",
        "]":"[",
    }

    stack := make([]string, 0)

    for i:= range parens{
        check := string(parens[i])
        if strings.Contains(closing,check) && len(stack) ==0{
            return false
        }else if strings.Contains(opening,check){
            stack = append(stack,check)
        }else if strings.Contains(closing,check){

            if dict[check] != stack[len(stack)-1]{
                return false
            }else if dict[check] == stack[len(stack)-1]{
                stack = stack[:len(stack)-1]
            }


        }
    }

    return len(stack) ==0

}
