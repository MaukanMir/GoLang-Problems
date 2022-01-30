Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

 

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]
 

Constraints:

1 <= n <= 8

Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Generate Parentheses.
Memory Usage: 2.8 MB, less than 92.75% of Go online submissions for Generate Parentheses.


func generateParenthesis(n int) []string {
    
    
    var result[]string
    
    backtrack("",0,0,&result,n)
    
    return result
}

func backtrack(check string, l int, r int, result *[]string, n int ){
    if len(check) == 2*n{
        *result = append(*result, check)
    }
    
    if(l < r){
        return
    }
    if(l<n){
        backtrack(check +"(", l+1,r, result, n)
    }
    
    if(l>r){
        backtrack(check +")", l,r+1, result,n)
    }
}
